package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
	"github.com/ruedigerp/newblog/internal/generator"
	"github.com/ruedigerp/newblog/internal/templates"
)

func Start(cfg *config.Config, posts []*content.Post, pages []*content.Page) {
	if err := templates.Init(cfg.ThemeDir()); err != nil {
		log.Fatalf("Failed to init templates: %v", err)
	}
	log.Printf("Using theme %q from %s", cfg.Theme, cfg.ThemeDir())

	tagMap := content.CollectTags(posts)
	catMap := content.CollectCategories(posts)
	postMap := make(map[string]*content.Post)
	postIndex := make(map[string]int)
	for i, p := range posts {
		postMap[p.Slug] = p
		postIndex[p.Slug] = i
	}
	pageMap := make(map[string]*content.Page)
	for _, p := range pages {
		pageMap[p.Slug] = p
	}

	site := templates.SiteData{
		Title:       cfg.Title,
		Description: cfg.Description,
		Author:      cfg.Author,
		BaseURL:     cfg.BaseURL,
		Menu:        cfg.Menu,
	}

	mux := http.NewServeMux()

	// Static files: the site's static/ wins, the theme's static/ is the fallback.
	// Mirrors the merge order used by the generator.
	mux.Handle("/static/", http.StripPrefix("/static/",
		staticHandler(cfg.StaticDir, cfg.ThemeStaticDir())))

	// Images (served from static/images, so /images/... paths in posts work directly)
	mux.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir(filepath.Join(cfg.StaticDir, "images")))))

	// RSS Feed
	mux.HandleFunc("/index.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
		rssCount := 20
		if rssCount > len(posts) {
			rssCount = len(posts)
		}
		if err := templates.RenderRSS(w, site, posts[:rssCount]); err != nil {
			log.Printf("Error rendering RSS: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// llms.txt — concise index for LLMs (https://llmstxt.org/)
	if cfg.LLMs.EnabledOrDefault() {
		mux.HandleFunc("/llms.txt", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
			if err := generator.RenderLLMsTxt(w, cfg, posts, pages); err != nil {
				log.Printf("Error rendering llms.txt: %v", err)
				http.Error(w, "Internal Server Error", 500)
			}
		})
	}

	// llms-full.txt — full markdown of every post
	if cfg.LLMs.FullOrDefault() {
		mux.HandleFunc("/llms-full.txt", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
			if err := generator.RenderLLMsFullTxt(w, cfg, posts); err != nil {
				log.Printf("Error rendering llms-full.txt: %v", err)
				http.Error(w, "Internal Server Error", 500)
			}
		})
	}

	// Search index
	mux.HandleFunc("/search-index.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		months := []string{
			"", "Januar", "Februar", "März", "April", "Mai", "Juni",
			"Juli", "August", "September", "Oktober", "November", "Dezember",
		}

		type searchEntry struct {
			Title      string   `json:"title"`
			Slug       string   `json:"slug"`
			Date       string   `json:"date"`
			Tags       []string `json:"tags"`
			Categories []string `json:"categories"`
			Preview    string   `json:"preview"`
			Cover      string   `json:"cover"`
			URL        string   `json:"url"`
		}

		entries := make([]searchEntry, 0, len(posts))
		for _, p := range posts {
			d := p.Date.Time
			dateStr := d.Format("02.") + " " + months[d.Month()] + " " + d.Format("2006")
			tags := p.Tags
			if tags == nil {
				tags = []string{}
			}
			cats := p.Categories
			if cats == nil {
				cats = []string{}
			}
			entries = append(entries, searchEntry{
				Title:      p.Title,
				Slug:       p.Slug,
				Date:       dateStr,
				Tags:       tags,
				Categories: cats,
				Preview:    p.GetPreview(),
				Cover:      p.GetCover(),
				URL:        "/posts/" + p.Slug + "/",
			})
		}

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		if err := enc.Encode(entries); err != nil {
			log.Printf("Error encoding search index: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// Pagination helper
	perPage := cfg.PostsPerPage
	totalPages := (len(posts) + perPage - 1) / perPage
	if totalPages < 1 {
		totalPages = 1
	}

	renderHomePage := func(w http.ResponseWriter, r *http.Request, page int) {
		if page < 1 || page > totalPages {
			http.NotFound(w, r)
			return
		}
		start := (page - 1) * perPage
		end := start + perPage
		if end > len(posts) {
			end = len(posts)
		}
		data := templates.HomeData{
			Site:       site,
			Posts:      posts[start:end],
			Page:       page,
			TotalPages: totalPages,
		}
		if err := templates.RenderHome(w, data); err != nil {
			log.Printf("Error rendering home: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	}

	// Homepage (page 1)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		renderHomePage(w, r, 1)
	})

	// Paginated pages: /page/2, /page/3, ...
	mux.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		numStr := strings.TrimPrefix(r.URL.Path, "/page/")
		numStr = strings.TrimSuffix(numStr, "/")
		page, err := strconv.Atoi(numStr)
		if err != nil || page < 1 {
			http.NotFound(w, r)
			return
		}
		if page == 1 {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		renderHomePage(w, r, page)
	})

	// Individual posts
	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/posts/")
		slug = strings.TrimSuffix(slug, "/")
		post, ok := postMap[slug]
		if !ok {
			http.NotFound(w, r)
			return
		}
		data := templates.PostData{Site: site, Post: post}
		// Posts are sorted newest first: index 0 = newest
		idx := postIndex[slug]
		if idx > 0 {
			data.PrevPost = posts[idx-1] // newer
		}
		if idx < len(posts)-1 {
			data.NextPost = posts[idx+1] // older
		}
		if err := templates.RenderPost(w, data); err != nil {
			log.Printf("Error rendering post: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// Static pages (sites)
	mux.HandleFunc("/sites/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/sites/")
		slug = strings.TrimSuffix(slug, "/")
		page, ok := pageMap[slug]
		if !ok {
			http.NotFound(w, r)
			return
		}
		data := templates.PageData{Site: site, Page: page}
		if err := templates.RenderPage(w, data); err != nil {
			log.Printf("Error rendering page: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// Tags index + single tag
	renderTagsIndex := func(w http.ResponseWriter) {
		tagCounts := make(map[string]int)
		for tag, tagPosts := range tagMap {
			tagCounts[tag] = len(tagPosts)
		}
		data := templates.TagsData{Site: site, Tags: tagCounts}
		if err := templates.RenderTags(w, data); err != nil {
			log.Printf("Error rendering tags: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	}
	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		renderTagsIndex(w)
	})
	mux.HandleFunc("/tags/", func(w http.ResponseWriter, r *http.Request) {
		tag := strings.TrimPrefix(r.URL.Path, "/tags/")
		tag = strings.TrimSuffix(tag, "/")
		if tag == "" {
			renderTagsIndex(w)
			return
		}
		handleTag(w, r, site, tagMap)
	})

	// Categories index + single category
	renderCategoriesIndex := func(w http.ResponseWriter) {
		catCounts := make(map[string]int)
		for cat, catPosts := range catMap {
			catCounts[cat] = len(catPosts)
		}
		data := templates.CategoriesData{Site: site, Categories: catCounts}
		if err := templates.RenderCategories(w, data); err != nil {
			log.Printf("Error rendering categories: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	}
	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		renderCategoriesIndex(w)
	})
	mux.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		cat := strings.TrimPrefix(r.URL.Path, "/categories/")
		cat = strings.TrimSuffix(cat, "/")
		if cat == "" {
			renderCategoriesIndex(w)
			return
		}
		handleCategory(w, r, site, catMap)
	})

	var handler http.Handler = mux
	if templates.DevMode() {
		// Templates are re-parsed per request; keep the browser from caching
		// the theme's CSS/JS so edits show up on reload.
		handler = noCache(mux)
		log.Printf("Dev mode: templates in %s are reloaded on every request", cfg.ThemeTemplatesDir())
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Blog server starting on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

// noCache disables browser caching — used in dev mode only.
func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, max-age=0")
		next.ServeHTTP(w, r)
	})
}

// staticHandler serves static assets from several directories: the first
// directory that contains the requested file wins.
func staticHandler(dirs ...string) http.Handler {
	servers := make([]http.Handler, len(dirs))
	for i, d := range dirs {
		servers[i] = http.FileServer(http.Dir(d))
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rel := filepath.FromSlash(path.Clean("/" + r.URL.Path))
		for i, d := range dirs {
			if info, err := os.Stat(filepath.Join(d, rel)); err == nil && !info.IsDir() {
				servers[i].ServeHTTP(w, r)
				return
			}
		}
		http.NotFound(w, r)
	})
}

func handleTag(w http.ResponseWriter, r *http.Request, site templates.SiteData, tagMap map[string][]*content.Post) {
	tag := strings.TrimPrefix(r.URL.Path, "/tags/")
	tag = strings.TrimSuffix(tag, "/")
	if tag == "" {
		http.Redirect(w, r, "/tags", http.StatusFound)
		return
	}
	for t, tagPosts := range tagMap {
		if strings.EqualFold(t, tag) {
			data := templates.TagData{Site: site, Tag: t, Posts: tagPosts}
			if err := templates.RenderTag(w, data); err != nil {
				log.Printf("Error rendering tag: %v", err)
				http.Error(w, "Internal Server Error", 500)
			}
			return
		}
	}
	http.NotFound(w, r)
}

func handleCategory(w http.ResponseWriter, r *http.Request, site templates.SiteData, catMap map[string][]*content.Post) {
	cat := strings.TrimPrefix(r.URL.Path, "/categories/")
	cat = strings.TrimSuffix(cat, "/")
	if cat == "" {
		http.Redirect(w, r, "/categories", http.StatusFound)
		return
	}
	for c, catPosts := range catMap {
		if strings.EqualFold(c, cat) {
			data := templates.CategoryData{Site: site, Category: c, Posts: catPosts}
			if err := templates.RenderCategory(w, data); err != nil {
				log.Printf("Error rendering category: %v", err)
				http.Error(w, "Internal Server Error", 500)
			}
			return
		}
	}
	http.NotFound(w, r)
}
