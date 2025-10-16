package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/PoulDev/lgBlog/internal/blog/model"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/ssh/terminal"
)

var DB *sql.DB
var bmp *bluemonday.Policy

func dbExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

func initSchema(db *sql.DB) {
	schema, err := os.ReadFile("./scripts/schema.sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	var passwd, pwdconfirm []byte
	var author model.Author;
	fmt.Println("Enter your name:")
	fmt.Scanln(&author.Name)

	fmt.Println("Enter your password:")
	passwd, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Confirm your password:")
	pwdconfirm, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	if string(passwd) != string(pwdconfirm) {
		log.Fatal("Passwords do not match")
	}

	author.Picture = "img/admin.jpg"

	_, err = CreateAccount(author, string(passwd), model.RoleAdmin)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Login using your username:", author.Name)

}

func LoadDB(filepath string) *sql.DB {
	bmp = bluemonday.UGCPolicy()

	exists := dbExists(filepath)

    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal(err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

	DB = db

	if !exists {
		log.Println("The database is new, initializing schema...")
		initSchema(db)
	} else {
		log.Println("The database already exists, skipping initialization")
	}

    return db
}
