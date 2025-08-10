package handlers

import (
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/PoulDev/lgBlog/internal/blog/config"
	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/golang-jwt/jwt"
)

type Message struct {
	model.BasePageData
	Type    string
	Message string
}

func loginPage(w http.ResponseWriter, r *http.Request, message Message) {
	fp := path.Join("web", "templates", "login.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loginApi(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		loginPage(w, r, Message{Type: "error", Message: "username and password are required"})
		return
	}

	account, err := db.Login(username, password)
	if err != nil {
		loginPage(w, r, Message{BasePageData: model.BasePageData{SiteTitle: config.Title, SiteDescription: config.Description, ShowCredits: config.ShowCredits, LoggedIn: false}, Type: "error", Message: err.Error()})
		return
	}

	// Set cookie
	tokenString, err := auth.GenToken(jwt.MapClaims{
		"uid": account.ID,
		"exp": time.Now().UTC().Add(time.Hour * 24 * 7).Unix(),
	})

	cookie := http.Cookie{Name: "token", Value: tokenString, HttpOnly: true, Secure: true, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		msg := Message{model.BasePageData{SiteTitle: config.Title, SiteDescription: config.Description, ShowCredits: config.ShowCredits, LoggedIn: false}, "info", "Please login to continue"}
		loginPage(w, r, msg)
	case http.MethodPost:
		loginApi(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
