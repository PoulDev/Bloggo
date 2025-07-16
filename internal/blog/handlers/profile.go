package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
	"github.com/PoulDev/lgBlog/internal/blog/model"
)

type Profile struct {
	model.Author
	Posts []model.Post
	PostsNum int
	IsItMe bool // is the client visiting the page the profile owner?
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

	// get token from cookie
	token, err := r.Cookie("token")
	isItMe := false
	if err == nil {
		claims, err := auth.CheckToken(token.Value)
		if err != nil {
			isItMe = false
		} else {
			
			isItMe = int64(claims["uid"].(float64)) == authorId
			log.Println(claims["uid"], authorId, isItMe)
		}
	}

	posts, err := db.GetPostsByAuthor(authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile := Profile{
		Author: author,
		IsItMe: isItMe,
		Posts: posts,
		PostsNum: len(posts),
	}

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
