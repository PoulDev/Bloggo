package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/PoulDev/lgBlog/internal/blog/config"
	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
	"github.com/PoulDev/lgBlog/internal/blog/model"
)

type WritePageData struct {
	model.BasePageData
	Post model.Post
	EditPost bool
}

func writePage(w http.ResponseWriter, r *http.Request, uid int64) {
	_, err := checkJWTcookie(r)
	loggedIn := err == nil

	pageData := WritePageData{
		BasePageData: model.BasePageData{
			SiteTitle: config.Title,
			SiteDescription: config.Description,
			LoggedIn: loggedIn,
		},
		Post: model.Post{},
		EditPost: false,
	}

	if r.URL.Query().Get("id") != "" {
		postId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pageData.EditPost = true 
		pageData.Post, err = db.GetPost(int64(postId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fp := path.Join("web", "templates", "write.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func writeApi(w http.ResponseWriter, r *http.Request, uid int64) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	description := r.FormValue("description")

	postId, err := db.NewPost(title, content, description, []int64{uid})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/post/%d", postId), http.StatusSeeOther)
}

func WriteHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if (err != nil) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, err := auth.CheckToken(token.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		writePage(w, r, int64(claims["uid"].(float64)))
	case http.MethodPost:
		writeApi(w, r, int64(claims["uid"].(float64)))
	}
}
