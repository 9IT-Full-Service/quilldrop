package generator

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
)

// RenderLLMsTxt writes an llms.txt index pointing LLMs at the site's resources.
// Follows the proposal at https://llmstxt.org/ — H1 title, blockquote summary,
// then H2 sections with markdown link lists.
func RenderLLMsTxt(w io.Writer, cfg *config.Config, posts []*content.Post, pages []*content.Page) error {
	bw := &errWriter{w: w}

	bw.printf("# %s\n\n", cfg.Title)
	bw.printf("> %s\n\n", cfg.Description)
	if cfg.Author != "" {
		bw.printf("Author: %s\n", cfg.Author)
	}
	bw.printf("Site: %s\n\n", cfg.BaseURL)
	bw.printf("Diese Datei folgt der llms.txt-Konvention (https://llmstxt.org/) und listet "+
		"alle Inhalte des Blogs, damit Sprachmodelle sie effizient finden und einordnen können. "+
		"Die vollständigen Texte liegen unter %s/llms-full.txt.\n\n", strings.TrimRight(cfg.BaseURL, "/"))

	// Blog posts
	bw.printf("## Blogposts\n\n")
	for _, p := range posts {
		url := strings.TrimRight(cfg.BaseURL, "/") + "/posts/" + p.Slug + "/"
		preview := strings.TrimSpace(p.GetPreview())
		date := p.Date.Format("2006-01-02")
		if preview != "" {
			bw.printf("- [%s](%s) — %s (%s)\n", p.Title, url, preview, date)
		} else {
			bw.printf("- [%s](%s) (%s)\n", p.Title, url, date)
		}
	}
	bw.printf("\n")

	// Static pages
	if len(pages) > 0 {
		bw.printf("## Seiten\n\n")
		for _, pg := range pages {
			url := strings.TrimRight(cfg.BaseURL, "/") + "/sites/" + pg.Slug + "/"
			if pg.Description != "" {
				bw.printf("- [%s](%s) — %s\n", pg.Title, url, pg.Description)
			} else {
				bw.printf("- [%s](%s)\n", pg.Title, url)
			}
		}
		bw.printf("\n")
	}

	// Optional resources
	base := strings.TrimRight(cfg.BaseURL, "/")
	bw.printf("## Optional\n\n")
	bw.printf("- [Tags](%s/tags/) — Übersicht aller Tags\n", base)
	bw.printf("- [Kategorien](%s/categories/) — Übersicht aller Kategorien\n", base)
	bw.printf("- [RSS Feed](%s/index.xml) — RSS-Feed mit den neuesten Posts\n", base)
	if cfg.LLMs.FullOrDefault() {
		bw.printf("- [Volltext](%s/llms-full.txt) — Alle Blogposts als Markdown in einer Datei\n", base)
	}

	return bw.err
}

// RenderLLMsFullTxt writes llms-full.txt — every post's full markdown content
// concatenated with metadata headers, suitable for LLM ingestion.
func RenderLLMsFullTxt(w io.Writer, cfg *config.Config, posts []*content.Post) error {
	bw := &errWriter{w: w}

	bw.printf("# %s\n\n", cfg.Title)
	bw.printf("> %s\n\n", cfg.Description)
	if cfg.Author != "" {
		bw.printf("Author: %s\n", cfg.Author)
	}
	bw.printf("Site: %s\n\n", cfg.BaseURL)
	if cfg.LLMs.EnabledOrDefault() {
		bw.printf("Dies ist der konsolidierte Volltext aller Blogposts (neueste zuerst). "+
			"Eine knappe Übersicht steht unter %s/llms.txt.\n\n", strings.TrimRight(cfg.BaseURL, "/"))
	} else {
		bw.printf("Dies ist der konsolidierte Volltext aller Blogposts (neueste zuerst).\n\n")
	}

	base := strings.TrimRight(cfg.BaseURL, "/")
	for i, p := range posts {
		if i > 0 {
			bw.printf("\n---\n\n")
		}
		url := base + "/posts/" + p.Slug + "/"
		bw.printf("## %s\n\n", p.Title)
		bw.printf("- URL: %s\n", url)
		bw.printf("- Datum: %s\n", p.Date.Format("2006-01-02"))
		if !p.Update.Time.IsZero() {
			bw.printf("- Aktualisiert: %s\n", p.Update.Format("2006-01-02"))
		}
		if p.Author != "" {
			bw.printf("- Autor: %s\n", p.Author)
		}
		if len(p.Tags) > 0 {
			bw.printf("- Tags: %s\n", strings.Join(p.Tags, ", "))
		}
		if len(p.Categories) > 0 {
			bw.printf("- Kategorien: %s\n", strings.Join(p.Categories, ", "))
		}
		bw.printf("\n")
		// Body is the preprocessed markdown (Hugo shortcodes already stripped).
		body := strings.TrimSpace(p.Content)
		bw.printf("%s\n", body)
	}

	return bw.err
}

// errWriter swallows repeated writes after the first error so callers don't
// have to check every Fprintf.
type errWriter struct {
	w   io.Writer
	err error
}

func (e *errWriter) printf(format string, args ...interface{}) {
	if e.err != nil {
		return
	}
	_, e.err = fmt.Fprintf(e.w, format, args...)
}

// GenerateLLMsFiles writes llms.txt and llms-full.txt into outDir,
// honoring cfg.LLMs.Enabled / cfg.LLMs.Full (both default to true).
func GenerateLLMsFiles(cfg *config.Config, posts []*content.Post, pages []*content.Page, outDir string) {
	if cfg.LLMs.EnabledOrDefault() {
		if err := writeFile(outDir+"/llms.txt", func(f *os.File) error {
			return RenderLLMsTxt(f, cfg, posts, pages)
		}); err != nil {
			log.Printf("Error writing llms.txt: %v", err)
		}
	}
	if cfg.LLMs.FullOrDefault() {
		if err := writeFile(outDir+"/llms-full.txt", func(f *os.File) error {
			return RenderLLMsFullTxt(f, cfg, posts)
		}); err != nil {
			log.Printf("Error writing llms-full.txt: %v", err)
		}
	}
}

func writeFile(path string, render func(f *os.File) error) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return render(f)
}
