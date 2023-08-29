package main

import (
	"flag"
)

type BuildConfig struct {
	dist string
	src  string
	css  string
	js   string
}

func InitFlagParser() BuildConfig {
	dist := flag.String("dist", "./dist", "build directory")
	src := flag.String("src", "./src", "source directory")

	css := flag.String("css", "./styles", "css file(s)")
	js := flag.String("js", "./scripts", "js file(s)")

	flag.Parse()
	c := BuildConfig{
		dist: *dist,
		src:  *src,
		css:  *css,
		js:   *js,
	}

	return c
}
