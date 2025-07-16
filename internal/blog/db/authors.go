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
	// This function doesn't return the post content to mantain it lightweight.
	// Why? because it's used only in /profile, and in that page only the post 
	// title and description are needed

	var posts []model.Post

	rows, err := DB.Query(`
        SELECT p.id, p.title, p.description
        FROM posts p
        JOIN post_authors pa ON pa.post = p.id
        WHERE pa.author = ?`, authorId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	for rows.Next() {
        var p model.Post
        if err := rows.Scan(&p.ID, &p.Title, &p.Description); err != nil {
            return nil, err
        }
        posts = append(posts, p)
    }

	return posts, nil
}
