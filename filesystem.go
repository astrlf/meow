package main

import (
	"os"

	"github.com/rs/zerolog/log"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func EnsureDir(path string) {
	if Exists(path) {
		return
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("failed to create directory")
	}
}
