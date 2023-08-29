package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	InitLogger()
	bc := InitFlagParser()

	log.Info().Msg("starting build")
	log.Info().
		Str("dist", bc.dist).
		Str("src", bc.src).
		Str("css", bc.css).
		Str("js", bc.js).
		Msg("build configuration")

	EnsureDir(bc.dist)
	if !Exists(bc.src) {
		log.Fatal().Msg("failed to create dist directory")
	}
}
