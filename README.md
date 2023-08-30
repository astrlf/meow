# 🐱 meow

a small static site generator for my personal website

# 🚀 usage

install from releases

```sh
$ tree posts
posts
├── main.md
├── first_post.md
└── second_post.md

$ meow
2023-08-31T00:52:31+02:00 INF main.go:16 > build configuration css=./style.css dist=./dist posts=./posts title="My Awesome Page"
2023-08-31T00:52:31+02:00 INF main.go:24 > starting build
2023-08-31T00:52:31+02:00 INF convert.go:70 > copied css file path=./style.css
2023-08-31T00:52:31+02:00 INF convert.go:114 > converting file file=posts/first_post.md
2023-08-31T00:52:31+02:00 INF convert.go:114 > converting file file=posts/second_post.md

$ tree dist
dist
├── index.html
├── posts/
│   ├── first_post.html
│   └── second_post.html
├── style.css
└── favicon.ico 
```

# 📝 notes

- meow defaults to `favicon.ico` in the `dist` directory.
- meow will make the `main.md` file in the `posts` directory the index page, and will move the blog posts to a `posts` directory in the `dist` directory.
- add serving
- add minification
- toc generation
