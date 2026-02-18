package content

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Page represents a static site page (e.g. Impressum, Datenschutz, Projekte).
type Page struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`

	// Derived
	Slug        string        // URL path relative to /sites/, e.g. "impressum" or "projekte/vm-tracker"
	Content     string        // raw markdown
	HTMLContent template.HTML // rendered HTML
}

// ParsePage reads a markdown file and returns a Page.
func ParsePage(path string, basePath string) (*Page, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := strings.TrimLeft(string(data), " \t\n\r")
	page := &Page{}

	if strings.HasPrefix(raw, "---") {
		// Has frontmatter
		rest := raw[3:]
		idx := strings.Index(rest, "\n---")
		if idx >= 0 {
			frontmatter := rest[:idx]
			body := rest[idx+4:]
			if err := yamlUnmarshal([]byte(frontmatter), page); err != nil {
				log.Printf("Warning: frontmatter parse error in %s: %v", path, err)
			}
			raw = body
		}
	}

	// Derive slug from relative path
	rel, _ := filepath.Rel(basePath, path)
	slug := strings.TrimSuffix(rel, filepath.Ext(rel))
	slug = strings.TrimSuffix(slug, ".md") // double .md handling
	// Normalize: if file is named index.md, use parent dir as slug
	if filepath.Base(slug) == "index" {
		slug = filepath.Dir(slug)
		if slug == "." {
			slug = ""
		}
	}
	page.Slug = slug

	// If no title from frontmatter, derive from slug
	if page.Title == "" {
		parts := strings.Split(slug, "/")
		last := parts[len(parts)-1]
		page.Title = strings.Title(strings.ReplaceAll(last, "-", " "))
	}

	// Render markdown
	rawContent := preprocessContent(raw)
	page.Content = rawContent

	var buf strings.Builder
	if err := mdConvert([]byte(rawContent), &buf); err != nil {
		log.Printf("Warning: markdown render error in %s: %v", path, err)
	}
	page.HTMLContent = template.HTML(buf.String())

	return page, nil
}

// LoadAllPages recursively reads all markdown files from the sites directory.
func LoadAllPages(sitesDir string) ([]*Page, error) {
	if _, err := os.Stat(sitesDir); os.IsNotExist(err) {
		return nil, nil // sites dir doesn't exist yet, that's ok
	}

	var pages []*Page
	err := filepath.WalkDir(sitesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		page, err := ParsePage(path, sitesDir)
		if err != nil {
			log.Printf("Warning: skipping page %s: %v", path, err)
			return nil
		}
		pages = append(pages, page)
		return nil
	})
	return pages, err
}
