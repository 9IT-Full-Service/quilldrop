package templates

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"
	"time"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
)

//go:embed *.html
var templateFS embed.FS

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
	homeTmpl       *template.Template
	postTmpl       *template.Template
	tagsTmpl       *template.Template
	tagTmpl        *template.Template
	categoriesTmpl *template.Template
	categoryTmpl   *template.Template
	pageTmpl       *template.Template
)

func parseTemplate(files ...string) (*template.Template, error) {
	return template.New("").Funcs(funcMap).ParseFS(templateFS, files...)
}

func Init() error {
	var err error
	homeTmpl, err = parseTemplate("base.html", "home.html")
	if err != nil {
		return fmt.Errorf("home template: %w", err)
	}
	postTmpl, err = parseTemplate("base.html", "post.html")
	if err != nil {
		return fmt.Errorf("post template: %w", err)
	}
	tagsTmpl, err = parseTemplate("base.html", "tags.html")
	if err != nil {
		return fmt.Errorf("tags template: %w", err)
	}
	tagTmpl, err = parseTemplate("base.html", "tag.html")
	if err != nil {
		return fmt.Errorf("tag template: %w", err)
	}
	categoriesTmpl, err = parseTemplate("base.html", "categories.html")
	if err != nil {
		return fmt.Errorf("categories template: %w", err)
	}
	categoryTmpl, err = parseTemplate("base.html", "category.html")
	if err != nil {
		return fmt.Errorf("category template: %w", err)
	}
	pageTmpl, err = parseTemplate("base.html", "page.html")
	if err != nil {
		return fmt.Errorf("page template: %w", err)
	}
	return nil
}

func RenderHome(w io.Writer, data HomeData) error {
	return homeTmpl.ExecuteTemplate(w, "base", data)
}

func RenderPost(w io.Writer, data PostData) error {
	return postTmpl.ExecuteTemplate(w, "base", data)
}

func RenderTags(w io.Writer, data TagsData) error {
	return tagsTmpl.ExecuteTemplate(w, "base", data)
}

func RenderTag(w io.Writer, data TagData) error {
	return tagTmpl.ExecuteTemplate(w, "base", data)
}

func RenderCategories(w io.Writer, data CategoriesData) error {
	return categoriesTmpl.ExecuteTemplate(w, "base", data)
}

func RenderCategory(w io.Writer, data CategoryData) error {
	return categoryTmpl.ExecuteTemplate(w, "base", data)
}

func RenderPage(w io.Writer, data PageData) error {
	return pageTmpl.ExecuteTemplate(w, "base", data)
}
