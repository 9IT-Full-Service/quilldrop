package generator

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
	"github.com/ruedigerp/newblog/internal/templates"
)

func Generate(cfg *config.Config, posts []*content.Post, pages []*content.Page) {
	if err := templates.Init(); err != nil {
		log.Fatalf("Failed to init templates: %v", err)
	}

	out := cfg.OutputDir

	// Clean and recreate output directory
	os.RemoveAll(out)
	os.MkdirAll(out, 0755)

	site := templates.SiteData{
		Title:       cfg.Title,
		Description: cfg.Description,
		Author:      cfg.Author,
		BaseURL:     cfg.BaseURL,
		Menu:        cfg.Menu,
	}

	// 1. Homepage with pagination
	perPage := cfg.PostsPerPage
	totalPages := (len(posts) + perPage - 1) / perPage
	if totalPages < 1 {
		totalPages = 1
	}
	for page := 1; page <= totalPages; page++ {
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
		if page == 1 {
			writeTemplate(filepath.Join(out, "index.html"), func(f *os.File) error {
				return templates.RenderHome(f, data)
			})
		}
		// Also generate /page/N/ for all pages (including page 1 as redirect target)
		if page > 1 {
			dir := filepath.Join(out, "page", fmt.Sprintf("%d", page))
			os.MkdirAll(dir, 0755)
			pageCopy := data // capture for closure
			writeTemplate(filepath.Join(dir, "index.html"), func(f *os.File) error {
				return templates.RenderHome(f, pageCopy)
			})
		}
	}

	// 2. Individual posts
	for i, post := range posts {
		dir := filepath.Join(out, "posts", post.Slug)
		os.MkdirAll(dir, 0755)
		data := templates.PostData{Site: site, Post: post}
		// Posts are sorted newest first: index 0 = newest
		if i > 0 {
			data.PrevPost = posts[i-1] // newer
		}
		if i < len(posts)-1 {
			data.NextPost = posts[i+1] // older
		}
		writeTemplate(filepath.Join(dir, "index.html"), func(f *os.File) error {
			return templates.RenderPost(f, data)
		})
	}

	// 3. Tags index
	tagMap := content.CollectTags(posts)
	tagCounts := make(map[string]int)
	for tag, tagPosts := range tagMap {
		tagCounts[tag] = len(tagPosts)
	}
	os.MkdirAll(filepath.Join(out, "tags"), 0755)
	writeTemplate(filepath.Join(out, "tags", "index.html"), func(f *os.File) error {
		return templates.RenderTags(f, templates.TagsData{Site: site, Tags: tagCounts})
	})

	// 4. Per-tag pages
	for tag, tagPosts := range tagMap {
		dir := filepath.Join(out, "tags", strings.ToLower(tag))
		os.MkdirAll(dir, 0755)
		writeTemplate(filepath.Join(dir, "index.html"), func(f *os.File) error {
			return templates.RenderTag(f, templates.TagData{Site: site, Tag: tag, Posts: tagPosts})
		})
	}

	// 5. Categories index
	catMap := content.CollectCategories(posts)
	catCounts := make(map[string]int)
	for cat, catPosts := range catMap {
		catCounts[cat] = len(catPosts)
	}
	os.MkdirAll(filepath.Join(out, "categories"), 0755)
	writeTemplate(filepath.Join(out, "categories", "index.html"), func(f *os.File) error {
		return templates.RenderCategories(f, templates.CategoriesData{Site: site, Categories: catCounts})
	})

	// 6. Per-category pages
	for cat, catPosts := range catMap {
		dir := filepath.Join(out, "categories", strings.ToLower(cat))
		os.MkdirAll(dir, 0755)
		writeTemplate(filepath.Join(dir, "index.html"), func(f *os.File) error {
			return templates.RenderCategory(f, templates.CategoryData{Site: site, Category: cat, Posts: catPosts})
		})
	}

	// 7. Static pages (sites)
	for _, page := range pages {
		dir := filepath.Join(out, "sites", page.Slug)
		os.MkdirAll(dir, 0755)
		writeTemplate(filepath.Join(dir, "index.html"), func(f *os.File) error {
			return templates.RenderPage(f, templates.PageData{Site: site, Page: page})
		})
	}

	// 8. RSS Feed
	rssCount := 20
	if rssCount > len(posts) {
		rssCount = len(posts)
	}
	writeTemplate(filepath.Join(out, "index.xml"), func(f *os.File) error {
		return templates.RenderRSS(f, site, posts[:rssCount])
	})

	// 9. Search index
	GenerateSearchIndex(posts, filepath.Join(out, "search-index.json"))

	// 10. Copy static assets (skip images/ subdirectory — handled separately)
	copyStaticDir("static", filepath.Join(out, "static"), []string{"images", "videos"})

	// 11. Copy images + videos to root /images/ and /videos/ (cover paths use /images/...)
	for _, dir := range []string{"images", "videos"} {
		src := filepath.Join("static", dir)
		dst := filepath.Join(out, dir)
		if _, err := os.Stat(src); err == nil {
			copyStaticDir(src, dst, nil)
		}
	}

	log.Printf("Generated %d posts, %d pages, %d tag pages, %d category pages, RSS feed into %s/",
		len(posts), len(pages), len(tagMap), len(catMap), out)
}

func writeTemplate(path string, render func(f *os.File) error) {
	f, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating %s: %v", path, err)
		return
	}
	defer f.Close()
	if err := render(f); err != nil {
		log.Printf("Error rendering %s: %v", path, err)
	}
}

func copyStaticDir(src, dst string, skipDirs []string) {
	skipSet := make(map[string]bool)
	for _, d := range skipDirs {
		skipSet[d] = true
	}
	filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		// Skip top-level directories in the skip list
		if d.IsDir() && rel != "." && skipSet[strings.Split(rel, string(filepath.Separator))[0]] {
			return filepath.SkipDir
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			os.MkdirAll(target, 0755)
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading %s: %v", path, err)
			return nil
		}
		os.MkdirAll(filepath.Dir(target), 0755)
		return os.WriteFile(target, data, 0644)
	})
}
