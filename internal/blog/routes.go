package blog

import (
	"net/http"

	"github.com/PoulDev/lgBlog/internal/blog/handlers"
	"github.com/microcosm-cc/bluemonday"
)

var (
	Bmp *bluemonday.Policy
)

func RegisterHandlers(mux *http.ServeMux) {
    mux.HandleFunc("/", handlers.Main)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/profile", handlers.ProfilePage)
	mux.HandleFunc("/write", handlers.WriteHandler)
	mux.HandleFunc("/post/", handlers.PostPage)

    mux.Handle("/css/", http.FileServer(http.Dir("./web/static/")))
}
