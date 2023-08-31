package main

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"

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

// file_name.md -> file_name
func RemoveExtension(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}

// file_name -> File Name
func NormaliseTitle(title string) string {
	words := strings.FieldsFunc(title, func(r rune) bool {
		return r == '_' || r == '-' || r == ' ' || r == '.'
	})

	for i, word := range words {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}
