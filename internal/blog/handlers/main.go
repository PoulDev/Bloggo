package handlers

import (
	"path"
	"net/http"
	"html/template"

	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/PoulDev/lgBlog/internal/blog/db"
)

type Posts struct {
	Posts []model.Post
	PostsNum int
}

func Main(w http.ResponseWriter, r *http.Request) {
	dbposts, err := db.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts := Posts{
		Posts: dbposts,
		PostsNum: len(dbposts),
	}

	fp := path.Join("web", "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
