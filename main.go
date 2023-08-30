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
		Str("title", bc.title).
		Msg("build configuration")

	if !Exists(bc.posts) {
		log.Fatal().
			Str("path", bc.posts).
			Msg("source directory does not exist")
	}

	log.Info().Msg("starting build")
	InitConvert(bc)
}
