package model

import (
	"time"
)

type Role int16

const (
	RoleAuthor Role = iota
	RoleAdmin
)

type Account struct {
	ID      int64
	Name    string
	Picture string
	Role    int16
	Pwd     []byte
	Salt    []byte
}

type Author struct {
	ID      int64
	Name    string
	Picture string
}

type Post struct {
	ID          int64
	Title       string
	Description string
	Content     string
	ContentRaw  string
	Created     time.Time
	Authors     []Author
}

type Comment struct {
	ID         int64
	Comment    string
	Created_at time.Time
	Post       int64
}
