package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
)

type SiteData struct {
	Title       string
	Description string
	Author      string
	BaseURL     string
	Menu        []config.MenuItem
}

type HomeData struct {
	Site       SiteData
	Posts      []*content.Post
	Page       int // current page (1-based)
	TotalPages int
}

type PostData struct {
	Site     SiteData
	Post     *content.Post
	PrevPost *content.Post // newer post (nil if this is the newest)
	NextPost *content.Post // older post (nil if this is the oldest)
}

type TagsData struct {
	Site SiteData
	Tags map[string]int
}

type TagData struct {
	Site  SiteData
	Tag   string
	Posts []*content.Post
}

type CategoriesData struct {
	Site       SiteData
	Categories map[string]int
}

type CategoryData struct {
	Site     SiteData
	Category string
	Posts    []*content.Post
}

type PageData struct {
	Site SiteData
	Page *content.Page
}

var funcMap = template.FuncMap{
	"formatDate": func(t time.Time) string {
		months := []string{
			"", "Januar", "Februar", "März", "April", "Mai", "Juni",
			"Juli", "August", "September", "Oktober", "November", "Dezember",
		}
		return t.Format("02.") + " " + months[t.Month()] + " " + t.Format("2006")
	},
	"isoDate": func(t time.Time) string {
		return t.Format("2006-01-02")
	},
	"rssDate": func(t time.Time) string {
		return t.Format(time.RFC1123Z)
	},
	"lower": strings.ToLower,
	"safeHTML": func(s template.HTML) template.HTML {
		return s
	},
	"hasChildren": func(item config.MenuItem) bool {
		return len(item.Children) > 0
	},
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"pageURL": func(page int) string {
		if page <= 1 {
			return "/"
		}
		return fmt.Sprintf("/page/%d/", page)
	},
	"seq": func(from, to int) []int {
		var s []int
		for i := from; i <= to; i++ {
			s = append(s, i)
		}
		return s
	},
	// paginationRange returns page numbers to display, with -1 as ellipsis placeholder.
	// Shows: first, last, and a window of 2 around the current page.
	"paginationRange": func(current, total int) []int {
		if total <= 7 {
			// Show all pages if 7 or fewer
			var s []int
			for i := 1; i <= total; i++ {
				s = append(s, i)
			}
			return s
		}
		pages := make(map[int]bool)
		pages[1] = true
		pages[total] = true
		for i := current - 2; i <= current+2; i++ {
			if i >= 1 && i <= total {
				pages[i] = true
			}
		}
		var result []int
		prev := 0
		for i := 1; i <= total; i++ {
			if pages[i] {
				if prev > 0 && i-prev > 1 {
					result = append(result, -1) // ellipsis
				}
				result = append(result, i)
				prev = i
			}
		}
		return result
	},
	"isEllipsis": func(n int) bool {
		return n == -1
	},
}

var (
	homeTmpl       *view
	postTmpl       *view
	tagsTmpl       *view
	tagTmpl        *view
	categoriesTmpl *view
	categoryTmpl   *view
	pageTmpl       *view
)

// templateDir holds the directory the templates of the active theme were
// loaded from (e.g. "themes/default/templates").
var templateDir string

// devMode re-parses templates on every render, so theme edits become visible
// without restarting the server.
var devMode bool

// SetDevMode enables or disables template live reload. Must be called before
// serving requests.
func SetDevMode(enabled bool) {
	devMode = enabled
}

// DevMode reports whether template live reload is enabled.
func DevMode() bool {
	return devMode
}

// view is a page template together with the theme files it was built from,
// so it can be re-parsed on demand in dev mode.
type view struct {
	files []string
	tmpl  *template.Template
}

func newView(files ...string) (*view, error) {
	tmpl, err := parseTemplate(files...)
	if err != nil {
		return nil, err
	}
	return &view{files: files, tmpl: tmpl}, nil
}

func (v *view) execute(w io.Writer, data any) error {
	if !devMode {
		return v.tmpl.ExecuteTemplate(w, "base", data)
	}
	// Live reload: parse the theme files again and render into a buffer first,
	// so a broken template does not emit half a page.
	tmpl, err := parseTemplate(v.files...)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "base", data); err != nil {
		return err
	}
	_, err = w.Write(buf.Bytes())
	return err
}

func parseTemplate(files ...string) (*template.Template, error) {
	paths := make([]string, 0, len(files))
	for _, f := range files {
		path := filepath.Join(templateDir, f)
		if _, err := os.Stat(path); err != nil {
			return nil, fmt.Errorf("missing template %s in theme: %w", f, err)
		}
		paths = append(paths, path)
	}
	return template.New("").Funcs(funcMap).ParseFiles(paths...)
}

// Init loads all templates of the theme located in themeDir
// (e.g. "themes/default"). The templates are expected in themeDir/templates.
func Init(themeDir string) error {
	templateDir = filepath.Join(themeDir, "templates")
	info, err := os.Stat(templateDir)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("theme templates not found in %s (check 'theme' and 'themesDir' in config.yaml)", templateDir)
	}

	homeTmpl, err = newView("base.html", "home.html")
	if err != nil {
		return fmt.Errorf("home template: %w", err)
	}
	postTmpl, err = newView("base.html", "post.html")
	if err != nil {
		return fmt.Errorf("post template: %w", err)
	}
	tagsTmpl, err = newView("base.html", "tags.html")
	if err != nil {
		return fmt.Errorf("tags template: %w", err)
	}
	tagTmpl, err = newView("base.html", "tag.html")
	if err != nil {
		return fmt.Errorf("tag template: %w", err)
	}
	categoriesTmpl, err = newView("base.html", "categories.html")
	if err != nil {
		return fmt.Errorf("categories template: %w", err)
	}
	categoryTmpl, err = newView("base.html", "category.html")
	if err != nil {
		return fmt.Errorf("category template: %w", err)
	}
	pageTmpl, err = newView("base.html", "page.html")
	if err != nil {
		return fmt.Errorf("page template: %w", err)
	}
	return nil
}

func RenderHome(w io.Writer, data HomeData) error {
	return homeTmpl.execute(w, data)
}

func RenderPost(w io.Writer, data PostData) error {
	return postTmpl.execute(w, data)
}

func RenderTags(w io.Writer, data TagsData) error {
	return tagsTmpl.execute(w, data)
}

func RenderTag(w io.Writer, data TagData) error {
	return tagTmpl.execute(w, data)
}

func RenderCategories(w io.Writer, data CategoriesData) error {
	return categoriesTmpl.execute(w, data)
}

func RenderCategory(w io.Writer, data CategoryData) error {
	return categoryTmpl.execute(w, data)
}

func RenderPage(w io.Writer, data PageData) error {
	return pageTmpl.execute(w, data)
}
