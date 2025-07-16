package main

import (
    "log"
    "net/http"

	"github.com/PoulDev/lgBlog/internal/blog"
	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db.LoadDB("blog.db")

    mux := http.NewServeMux()

    blog.RegisterHandlers(mux)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
