package main

import (
	"flag"
)

type BuildConfig struct {
	posts string
	dist  string
	css   string
}

func InitFlagParser() BuildConfig {
	posts := flag.String("posts", "./posts", "posts directory")
	dist := flag.String("dist", "./dist", "build directory")
	css := flag.String("css", "./style.css", "css file")

	flag.Parse()
	c := BuildConfig{
		dist:  *dist,
		posts: *posts,
		css:   *css,
	}

	return c
}
