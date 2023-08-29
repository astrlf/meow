package main

import "github.com/rs/zerolog/log"

func main() {
	InitLogger()

	log.Info().Str("foo", "bar").Msg("Hello world")
}
