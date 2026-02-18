package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ruedigerp/newblog/internal/config"
	"github.com/ruedigerp/newblog/internal/content"
	"github.com/ruedigerp/newblog/internal/generator"
	"github.com/ruedigerp/newblog/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: newblog <serve|generate>")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  serve      Start the dynamic HTTP server")
		fmt.Println("  generate   Generate static HTML files")
		os.Exit(1)
	}

	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
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

	switch os.Args[1] {
	case "serve":
		server.Start(cfg, posts, pages)
	case "generate":
		generator.Generate(cfg, posts, pages)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Usage: newblog <serve|generate>")
		os.Exit(1)
	}
}
