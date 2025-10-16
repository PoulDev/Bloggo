package blog

import (
	"net/http"

	"github.com/PoulDev/lgBlog/internal/blog/config"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
	"github.com/PoulDev/lgBlog/internal/blog/handlers"
	"github.com/microcosm-cc/bluemonday"
)

var (
	Bmp *bluemonday.Policy
)

type ServeMux struct {
	*http.ServeMux
}

func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if !config.PrivateBlog {
		mux.ServeMux.HandleFunc(pattern, handler)
		return
	}

	mux.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if (err != nil) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		_, err = auth.CheckToken(token.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		handler(w, r)
	})
}

func RegisterHandlers(mux ServeMux) {
    mux.HandleFunc("/", handlers.Main)
	mux.HandleFunc("/profile", handlers.ProfilePage)
	mux.HandleFunc("/edit", handlers.EditProfilePage)
	mux.HandleFunc("/write", handlers.WriteHandler)
	mux.HandleFunc("/post/", handlers.PostPage)

	mux.ServeMux.HandleFunc("/login", handlers.Login)

    mux.Handle("/css/", http.FileServer(http.Dir("./web/static/")))
    mux.Handle("/img/", http.FileServer(http.Dir("./web/static/")))
}
