package handlers

import (
	"html/template"
	"net/http"
	"path"

	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/PoulDev/lgBlog/internal/blog/config"
)

type MainPage struct {
	model.BasePageData

	Posts []model.Post
	PostsNum int
	LoggedIn bool
}

func Main(w http.ResponseWriter, r *http.Request) {
	_, err := checkJWTcookie(r)
	loggedIn := err == nil

	dbposts, err := db.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	maindata := MainPage{
		BasePageData: model.BasePageData{SiteTitle: config.Title, SiteDescription: config.Description, LoggedIn: loggedIn},
		Posts: dbposts,
		PostsNum: len(dbposts),
		LoggedIn: loggedIn,
	}

	fp := path.Join("web", "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, maindata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
