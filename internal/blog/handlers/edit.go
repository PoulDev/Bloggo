package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PoulDev/lgBlog/internal/blog/db"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
)

func editPostApi(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if (err != nil) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	_, err = auth.CheckToken(token.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	description := r.FormValue("description")

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	postid, err := strconv.Atoi(parts[1])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = db.UpdatePost(int64(postid), title, content, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/post/%d", postid), http.StatusSeeOther)
}

