package main

import (
	"flag"
)

type BuildConfig struct {
	posts string
	dist  string
	css   string
	title string
}

func InitFlagParser() BuildConfig {
	posts := flag.String("posts", "./posts", "posts directory")
	dist := flag.String("dist", "./dist", "build directory")
	css := flag.String("css", "./style.css", "css file")
	title := flag.String("title", "My Awesome Page", "website name")

	flag.Parse()
	c := BuildConfig{
		dist:  *dist,
		posts: *posts,
		css:   *css,
		title: *title,
	}

	return c
}
