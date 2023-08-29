package main

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func ConvertToHTML(path string) {
	files, err := filepath.Glob(filepath.Join(path, "*.md"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to glob path")
	}

	for _, file := range files {
		log.Info().Str("file", file).Msg("found file")
	}
}
