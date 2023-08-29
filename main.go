package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	InitLogger()
	bc := InitFlagParser()

	log.Info().
		Str("posts", bc.posts).
		Str("dist", bc.dist).
		Str("css", bc.css).
		Str("js", bc.js).
		Msg("build configuration")

	if !Exists(bc.posts) {
		log.Fatal().Msg("source directory does not exist")
	}

	log.Info().Msg("starting build")
	ConvertToHTML(bc.posts)
}
