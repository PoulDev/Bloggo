package handlers

import (
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/PoulDev/lgBlog/internal/blog/config"
)

func EditProfilePage(w http.ResponseWriter, r *http.Request) {
	authorIDstr := r.URL.Query().Get("author")
	
	// get token from cookie
	loggedUserId := int64(0)
	isItMe := false

	token, err := r.Cookie("token")
	if err == nil {
		claims, err := auth.CheckToken(token.Value)
		if err == nil {
			loggedUserId = int64(claims["uid"].(float64))
		}
	}

	authorId, err := strconv.ParseInt(authorIDstr, 10, 64)
	if err != nil {
		if loggedUserId != 0 {
			authorId = loggedUserId
		} else {
			http.Error(w, "Invalid author ID!", http.StatusInternalServerError)
			return
		}
	}

	if loggedUserId != 0 {
		isItMe = loggedUserId == authorId
	}

	author, err := db.GetAuthor(authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	posts, err := db.GetPostsByAuthor(authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile := Profile{
		Author: author,
		BasePageData: model.BasePageData{SiteTitle: config.Title, SiteDescription: config.Description, ShowCredits: config.ShowCredits, LoggedIn: loggedUserId != 0},

		IsItMe: isItMe,
		Posts: posts,
		PostsNum: len(posts),
	}

	fp := path.Join("web", "templates", "editprofile.html")
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
