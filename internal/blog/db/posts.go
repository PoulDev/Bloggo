package db

import (
	"strings"
	"time"

	"github.com/PoulDev/lgBlog/internal/blog/model"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(content string) []byte {
	content = strings.ReplaceAll(content, "\n\n", "\n\n<br>\n\n")
	content = strings.ReplaceAll(content, "\n\r\n\r", "\n\n<br>\n\n")
	content = strings.ReplaceAll(content, "\r\n\r\n", "\n\n<br>\n\n")

	md := []byte(content)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.Footnotes | parser.HardLineBreak
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func GetPosts() ([]model.Post, error) {
	var posts []model.Post

	rows, err := DB.Query(`
        SELECT p.id, p.title, p.description, p.content, p.contentRaw, p.created_at
        FROM posts p
        ORDER BY p.created_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	for rows.Next() {
        var p model.Post
        if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Content, &p.ContentRaw, &p.Created); err != nil {
            return nil, err
        }
        posts = append(posts, p)
    }

	return posts, nil
}

func GetPost(postId int64) (model.Post, error) {
	var p model.Post
	err := DB.QueryRow("SELECT id, title, description, content, contentRaw, created_at FROM posts WHERE id = ?", postId).Scan(&p.ID, &p.Title, &p.Description, &p.Content, &p.ContentRaw, &p.Created)
	if err != nil {
		return p, err
	}
	return p, nil
}

func GetAuthors(postId int64) ([]model.Author, error) {
	var authors []model.Author

	
	rows, err := DB.Query(`
        SELECT a.id, a.name, a.picture
        FROM authors a
        JOIN post_authors pa ON pa.author = a.id
        WHERE pa.post = ?`, postId)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a model.Author
		if err := rows.Scan(&a.ID, &a.Name, &a.Picture); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	return authors, nil
}

func NewPost(title string, content string, description string, authors []int64) (int64, error) {
	markdown := string(mdToHTML(content))
	sanitized := bmp.Sanitize(markdown)

	result, err := DB.Exec("INSERT INTO posts (title, description, content, contentRaw, created_at) VALUES (?, ?, ?, ?, ?)", title, description, sanitized, content, time.Now().UTC().Unix())

	if err != nil {
		return -1, err
	}

	post, err := result.LastInsertId()
	if (err != nil) {
		return -1, err
	}

	for _, authorId := range authors {
		_, err := DB.Exec("INSERT INTO post_authors (post, author) VALUES (?, ?)", post, authorId)
		if err != nil {
			return -1, err
		}
	}

    return post, nil
}

func DeletePost(id int64) error {
	_, err := DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}

	_, err = DB.Exec("DELETE FROM post_authors WHERE post = ?", id)
	if err != nil {
		return err
	}

    return nil
}

func UpdatePost(id int64, title string, content string, description string) error {
	markdown := string(mdToHTML(content))
	sanitized := bmp.Sanitize(markdown)

	_, err := DB.Exec("UPDATE posts SET title = ?, description = ?, content = ?, contentRaw = ? WHERE id = ?", title, description, sanitized, content, id)
	if err != nil {
		return err
	}

	return nil
}
