package db;

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/PoulDev/lgBlog/internal/blog/db/auth"
)

func RandomPassword() (string, error) {
    b := make([]byte, 32)

    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return "", err
    }

    return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b), nil
}

func CreateAccount(author model.Author, password string, role model.Role) (int64, error) {
	hash, salt, err := auth.HashPassword(password)
	if err != nil {
		return -1, err
	}

	result, err := DB.Exec("INSERT INTO authors (name, picture, role, pwd, salt) VALUES (?, ?, ?, ?, ?)",
		author.Name, author.Picture, role, hash, salt)
	
	if err != nil {
		return -1, err
	}

    id, err := result.LastInsertId()
    if err != nil {
        return -1, err
    }

	return id, nil;
}

func Login(username string, password string) (model.Account, error) {
	var account model.Account

	err := DB.QueryRow("SELECT id, name, picture, role, pwd, salt FROM authors WHERE name = ?", username).
		Scan(&account.ID, &account.Name, &account.Picture, &account.Role, &account.Pwd, &account.Salt)
	if err != nil {
		return account, err
	}

	if !auth.CheckPassword(password, account.Pwd, account.Salt) {
		return model.Account{}, errors.New("invalid password")
	}

	return account, nil
}
