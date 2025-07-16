package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"strconv"
	"path"
)

type Post struct {
	Title string
	Content string
}

func PostPage(w http.ResponseWriter, r *http.Request) {
	//author := r.URL.Query().Get("author")
	// TODO: Get author from database
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(parts) != 2 || parts[0] != "post" {
        http.NotFound(w, r)
        return
    }

	_, err := strconv.Atoi(parts[1])
	if err != nil {
		http.NotFound(w, r)
		return
	}


	post := Post{
		Title: "Hello, world!",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor...",
	}

	fp := path.Join("web", "templates", "post.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
