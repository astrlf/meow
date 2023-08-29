package main

import (
	"flag"
)

type BuildConfig struct {
	posts string
	dist  string
	css   string
	js    string
}

func InitFlagParser() BuildConfig {
	posts := flag.String("posts", "./posts", "posts directory")
	dist := flag.String("dist", "./dist", "build directory")
	css := flag.String("css", "./styles", "css file(s)")
	js := flag.String("js", "./scripts", "js file(s)")

	flag.Parse()
	c := BuildConfig{
		dist:  *dist,
		posts: *posts,
		css:   *css,
		js:    *js,
	}

	return c
}
