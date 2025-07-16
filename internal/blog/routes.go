package blog

import (
	"net/http"

	"github.com/PoulDev/lgBlog/internal/blog/handlers"
)

func RegisterHandlers(mux *http.ServeMux) {
    mux.Handle("/", http.FileServer(http.Dir("./web/static/")))
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/profile", handlers.ProfilePage)

	// handle /post/{id}
	mux.HandleFunc("/post/", handlers.PostPage)
}
