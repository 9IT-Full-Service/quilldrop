package generator

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/ruedigerp/newblog/internal/config"
)

// RenderRobotsTxt writes a robots.txt for the site. Without explicit rules the
// whole site is allowed; the sitemap is referenced automatically whenever
// sitemap generation is enabled.
func RenderRobotsTxt(w io.Writer, cfg *config.Config) error {
	bw := &errWriter{w: w}

	bw.printf("User-agent: %s\n", cfg.Robots.UserAgentOrDefault())
	for _, rule := range cfg.Robots.Allow {
		bw.printf("Allow: %s\n", rule)
	}
	for _, rule := range cfg.Robots.Disallow {
		bw.printf("Disallow: %s\n", rule)
	}
	// A crawler needs at least one rule; allow everything when none are configured.
	if len(cfg.Robots.Allow) == 0 && len(cfg.Robots.Disallow) == 0 {
		bw.printf("Allow: /\n")
	}

	if cfg.Sitemap.EnabledOrDefault() && cfg.BaseURL != "" {
		bw.printf("\nSitemap: %s/sitemap.xml\n", strings.TrimRight(cfg.BaseURL, "/"))
	}

	return bw.err
}

// GenerateRobotsTxt writes robots.txt into outDir, honoring cfg.Robots.Enabled
// (defaults to true).
func GenerateRobotsTxt(cfg *config.Config, outDir string) {
	if !cfg.Robots.EnabledOrDefault() {
		return
	}
	if err := writeFile(outDir+"/robots.txt", func(f *os.File) error {
		return RenderRobotsTxt(f, cfg)
	}); err != nil {
		log.Printf("Error writing robots.txt: %v", err)
	}
}
