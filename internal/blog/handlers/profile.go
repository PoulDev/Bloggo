package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"path"

	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/PoulDev/lgBlog/internal/blog/db"
)

type Profile struct {
	model.Author
	Posts []model.Post
	PostsNum int
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	authorIDstr := r.URL.Query().Get("author")
	// TODO: Get author from database

	authorId, err := strconv.ParseInt(authorIDstr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid author ID!", http.StatusInternalServerError)
		return
	}

	author, err := db.GetAuthor(authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile := Profile{
		Author: author,
		Posts: []model.Post{
			{
				ID: 1,
				Title: "Hello, world!",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor...",
			},
			{
				ID: 2,
				Title: "Hello, world!",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor...",
			},
			{
				ID: 3,
				Title: "Hello, world!",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor...",
			},
			{
				ID: 4,
				Title: "Hello, world!",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor...",
			},
			{
				ID: 5,
				Title: "Hello, world!",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor...",
			},
		},
	}
	profile.PostsNum = len(profile.Posts)

	fp := path.Join("web", "templates", "author.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
