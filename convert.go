package main

import (
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/rs/zerolog/log"
)

func ConvertToHTML(md []byte) []byte {
	document := parser.
		NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock).
		Parse(md)
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank})

	return markdown.Render(document, renderer)
}

func InitConvert(config BuildConfig) {
	files, err := filepath.Glob(filepath.Join(config.posts, "*.md"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to glob path")
	}

	EnsureDir(config.dist)
	for _, file := range files {
		log.Info().Str("file", file).Msg("converting file")

		md, err := os.ReadFile(file)
		if err != nil {
			log.Fatal().Err(err).Str("file", file).Msg("failed to read file")
		}

		html := ConvertToHTML(md)

		htmlFile := filepath.Join(config.dist, filepath.Base(file))
		htmlFile = htmlFile[:len(htmlFile)-len(filepath.Ext(htmlFile))] + ".html"

		if err := os.WriteFile(htmlFile, html, 0644); err != nil {
			log.Fatal().Err(err).Str("file", htmlFile).Msg("failed to write file")
		}
	}
}
