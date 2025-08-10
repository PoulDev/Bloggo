package handlers

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/PoulDev/lgBlog/internal/blog/config"
	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/model"
)

type Post struct {
	model.Post
	model.BasePageData

	Content template.HTML
}

func PostPage(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

    if len(parts) == 3 && parts[0] == "post" {
		if parts[2] == "delete" {
			deletePostApi(w, r)
		} else if parts[2] == "edit" {
			editPostApi(w, r)
		}
        http.NotFound(w, r)
		return
	}

	if len(parts) != 2 {
        http.NotFound(w, r)
        return
    }

	postid, err := strconv.Atoi(parts[1])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	_, err = checkJWTcookie(r)
	loggedIn :=  err == nil


	post, err := db.GetPost(int64(postid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.Authors, err = db.GetAuthors(post.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := Post{post, model.BasePageData{SiteTitle: config.Title, SiteDescription: config.Description, ShowCredits: config.ShowCredits, LoggedIn: loggedIn}, template.HTML(string(post.Content))}

	fp := path.Join("web", "templates", "post.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
