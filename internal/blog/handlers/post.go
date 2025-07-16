package handlers

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/model"
)

type Post struct {
	model.Post
	Content template.HTML
}

func PostPage(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(parts) != 2 || parts[0] != "post" {
        http.NotFound(w, r)
        return
    }

	postid, err := strconv.Atoi(parts[1])
	if err != nil {
		http.NotFound(w, r)
		return
	}


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

	data := Post{post, template.HTML(string(post.Content))}

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

