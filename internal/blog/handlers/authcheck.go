package handlers

import (
	"net/http"

	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
)

func checkJWTcookie(r *http.Request) (int64, error) {
	token, err := r.Cookie("token")
	if (err != nil) {
		return 0, err
	}
	claims, err := auth.CheckToken(token.Value)
	if err != nil {
		return 0, err
	}

	return int64(claims["uid"].(float64)), nil
}
