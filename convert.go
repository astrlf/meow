package main

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/rs/zerolog/log"
)

const htmlTemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>{{ .Title }}</title>
<link rel="stylesheet" href="{{ .CSS }}">
<link rel="icon" href="favicon.ico">
</head>

<body>
<h1 class="meow__title">{{ .Title }}</h1>
<div class="meow__container">{{ .Content }}</div>
</body>
</html>
`

type TemplateVariables struct {
	Title   string
	CSS     string
	Content string
}

func TemplateFile(config BuildConfig, html []byte) []byte {
	var templated bytes.Buffer

	tmpl := template.Must(template.New("html").Parse(htmlTemplate))
	err := tmpl.Execute(&templated, TemplateVariables{
		CSS:     config.css,
		Title:   config.title,
		Content: string(html),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("failed to execute template")
	}

	return templated.Bytes()
}

func EnsureCSS(config *BuildConfig) {
	content, err := os.ReadFile(config.css)
	if err != nil {
		// sane default
		config.css = "https://unpkg.com/sakura.css/css/sakura.css"
		log.Warn().
			Str("path", config.css).
			Msg("could not read css file, defaulting to sakura.css")

		return
	}

	os.WriteFile(filepath.Join(config.dist, "style.css"), content, 0644)
	log.Info().
		Str("path", config.css).
		Msg("copied css file")
}

func EnsureFavicon(config BuildConfig) {
	favicon := "favicon.ico"
	content, err := os.ReadFile(favicon)
	if err != nil {
		log.Warn().
			Str("path", favicon).
			Msg("could not read favicon file, defaulting to none")

		return
	}

	os.WriteFile(filepath.Join(config.dist, "favicon.ico"), content, 0644)
	log.Info().
		Str("path", favicon).
		Msg("copied favicon")
}

func EnsureTOC(config BuildConfig) {
	return // TODO
}

func EnsureMain(config BuildConfig) {
	main := filepath.Join(config.posts, "main.md")
	if _, err := os.Stat(main); os.IsNotExist(err) {
		log.Fatal().Err(err).Msg("main.md does not exist")
	}

	md, err := os.ReadFile(main)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read main.md")
	}

	os.WriteFile(filepath.Join(config.dist, "index.html"), ConvertToHTML(config, md), 0644)
}

func ConvertToHTML(config BuildConfig, md []byte) []byte {
	document := parser.
		NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock).
		Parse(md)
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank})

	return TemplateFile(config, markdown.Render(document, renderer))
}

func InitConvert(config BuildConfig) {
	files, err := filepath.Glob(filepath.Join(config.posts, "*.md"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to glob path")
	}

	EnsureDir(config.dist)
	EnsureDir(filepath.Join(config.dist, "posts"))

	EnsureCSS(&config)
	EnsureMain(config)
	EnsureFavicon(config)

	for _, file := range files {
		if filepath.Base(file) == "main.md" {
			continue
		}

		log.Info().Str("file", file).Msg("converting file")

		md, err := os.ReadFile(file)
		if err != nil {
			log.Fatal().Err(err).Str("file", file).Msg("failed to read file")
		}

		config.title = NormaliseTitle(RemoveExtension(filepath.Base(file)))

		htmlFile := filepath.Join(config.dist, "posts", filepath.Base(file))
		htmlFile = RemoveExtension(htmlFile) + ".html"

		if err := os.WriteFile(htmlFile, ConvertToHTML(config, md), 0644); err != nil {
			log.Fatal().Err(err).Str("file", htmlFile).Msg("failed to write file")
		}
	}
}
