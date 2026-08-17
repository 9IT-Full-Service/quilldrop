package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
	"github.com/ruedigerp/newblog/internal/generator"
	"github.com/ruedigerp/newblog/internal/server"
	"github.com/ruedigerp/newblog/internal/templates"
)

func usage() {
	fmt.Println("Usage: quilldrop <serve|generate> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  serve      Start the dynamic HTTP server")
	fmt.Println("  generate   Generate static HTML files")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -config <path>   Config file (default: config.yaml)")
	fmt.Println("  -theme <name>    Theme to use, overrides 'theme' from the config")
	fmt.Println("  -dev             serve: reload templates on every request (live reload)")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	if cmd != "serve" && cmd != "generate" {
		fmt.Printf("Unknown command: %s\n", cmd)
		usage()
		os.Exit(1)
	}

	fs := flag.NewFlagSet(cmd, flag.ExitOnError)
	configPath := fs.String("config", "config.yaml", "path to the config file")
	theme := fs.String("theme", "", "theme to use, overrides 'theme' from the config")
	dev := fs.Bool("dev", false, "serve: reload templates on every request")
	if err := fs.Parse(os.Args[2:]); err != nil {
		os.Exit(1)
	}
	if *dev && cmd != "serve" {
		log.Fatalf("-dev is only supported for 'serve'")
	}
	templates.SetDevMode(*dev)

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	if *theme != "" {
		cfg.Theme = *theme
	}
	if err := checkTheme(cfg); err != nil {
		log.Fatalf("%v", err)
	}

	posts, err := content.LoadAll(cfg.ContentDir)
	if err != nil {
		log.Fatalf("Failed to load posts: %v", err)
	}
	log.Printf("Loaded %d posts", len(posts))

	pages, err := content.LoadAllPages(cfg.SitesDir)
	if err != nil {
		log.Fatalf("Failed to load pages: %v", err)
	}
	log.Printf("Loaded %d pages", len(pages))

	switch cmd {
	case "serve":
		server.Start(cfg, posts, pages)
	case "generate":
		generator.Generate(cfg, posts, pages)
	}
}

// checkTheme verifies the configured theme exists and reports the available
// themes otherwise.
func checkTheme(cfg *config.Config) error {
	if info, err := os.Stat(cfg.ThemeTemplatesDir()); err == nil && info.IsDir() {
		return nil
	}
	msg := fmt.Sprintf("Theme %q not found: %s is missing", cfg.Theme, cfg.ThemeTemplatesDir())
	if available := availableThemes(cfg.ThemesDir); len(available) > 0 {
		msg += fmt.Sprintf(" (available themes: %s)", strings.Join(available, ", "))
	}
	return fmt.Errorf("%s", msg)
}

func availableThemes(themesDir string) []string {
	entries, err := os.ReadDir(themesDir)
	if err != nil {
		return nil
	}
	var themes []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if info, err := os.Stat(filepath.Join(themesDir, e.Name(), "templates")); err == nil && info.IsDir() {
			themes = append(themes, e.Name())
		}
	}
	sort.Strings(themes)
	return themes
}
