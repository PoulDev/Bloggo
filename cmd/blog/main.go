package main

import (
	"fmt"
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

	log.Println("Starting server on port", config.HostPort)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", config.HostPort), mux); err != nil {
        log.Fatal(err)
    }
}
