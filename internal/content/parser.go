package content

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

var md goldmark.Markdown

// yamlUnmarshal is a package-level wrapper for yaml.Unmarshal.
var yamlUnmarshal = yaml.Unmarshal

func init() {
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
				highlighting.WithFormatOptions(
					chromahtml.WithClasses(true),
				),
			),
			emoji.Emoji,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
}

var hugoShortcodeRe = regexp.MustCompile(`\{\{<\s*/?[a-zA-Z][^>]*>\}\}`)

// mdConvert renders markdown to the given writer using the shared goldmark instance.
func mdConvert(source []byte, w *strings.Builder) error {
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return err
	}
	w.WriteString(buf.String())
	return nil
}

// preprocessContent strips Hugo shortcode delimiters while preserving raw HTML content between them.
func preprocessContent(raw string) string {
	// Remove {{< rawhtml >}} and {{< /rawhtml >}} delimiters
	raw = strings.ReplaceAll(raw, "{{< rawhtml >}}", "")
	raw = strings.ReplaceAll(raw, "{{< /rawhtml >}}", "")
	// Remove any remaining Hugo shortcodes
	raw = hugoShortcodeRe.ReplaceAllString(raw, "")
	return raw
}

// ParseFile reads a markdown file and returns a Post with parsed frontmatter and rendered HTML.
func ParseFile(path string) (*Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Split on --- frontmatter delimiters
	// Trim leading whitespace/newlines (some files have a newline before ---)
	content := strings.TrimLeft(string(data), " \t\n\r")
	if !strings.HasPrefix(content, "---") {
		return nil, fmt.Errorf("no frontmatter found in %s", path)
	}

	// Find the closing ---
	rest := content[3:]
	idx := strings.Index(rest, "\n---")
	if idx < 0 {
		return nil, fmt.Errorf("unclosed frontmatter in %s", path)
	}

	frontmatter := rest[:idx]
	body := rest[idx+4:] // skip past \n---

	post := &Post{}
	if err := yaml.Unmarshal([]byte(frontmatter), post); err != nil {
		return nil, fmt.Errorf("parse frontmatter %s: %w", path, err)
	}

	// Derive slug from filename
	base := filepath.Base(path)
	post.Filename = base
	slug := strings.TrimSuffix(base, filepath.Ext(base))
	// Handle double .md extension
	slug = strings.TrimSuffix(slug, ".md")
	post.Slug = slug

	// Preprocess and render markdown
	rawContent := preprocessContent(body)
	post.Content = rawContent

	var buf bytes.Buffer
	if err := md.Convert([]byte(rawContent), &buf); err != nil {
		return nil, fmt.Errorf("render markdown %s: %w", path, err)
	}
	post.HTMLContent = template.HTML(buf.String())

	return post, nil
}

// LoadAll reads all markdown files from a directory, filters out drafts/hidden, and sorts by date descending.
func LoadAll(contentDir string) ([]*Post, error) {
	entries, err := os.ReadDir(contentDir)
	if err != nil {
		return nil, err
	}

	var posts []*Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		post, err := ParseFile(filepath.Join(contentDir, entry.Name()))
		if err != nil {
			log.Printf("Warning: skipping %s: %v", entry.Name(), err)
			continue
		}
		if post.Draft || post.Hide {
			continue
		}
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date.Time)
	})

	return posts, nil
}
