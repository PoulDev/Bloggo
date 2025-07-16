package db;

import (
	"github.com/PoulDev/lgBlog/internal/blog/model"
)

func GetAuthor(authorId int64) (model.Author, error) {
	var author model.Author

	err := DB.QueryRow("SELECT id, name, picture FROM authors WHERE id = ?", authorId).
		Scan(&author.ID, &author.Name, &author.Picture)

	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func GetPostsByAuthor(authorId int64) ([]model.Post, error) {
	var posts []model.Post

	rows, err := DB.Query(`
        SELECT p.id, p.title, p.content, p.created_at
        FROM posts p
        JOIN post_authors pa ON pa.post = p.id
        WHERE pa.author = ?`, authorId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	for rows.Next() {
        var p model.Post
        if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Created); err != nil {
            return nil, err
        }
        posts = append(posts, p)
    }

	return posts, nil
}
