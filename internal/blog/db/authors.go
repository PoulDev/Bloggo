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
