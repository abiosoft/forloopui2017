package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Post struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func createPost(p Post) error {
	_, err := db.Exec("INSERT INTO guestbook (guestName, content) values ($1, $2)", p.Name, p.Content)
	return err
}

func getPosts() ([]Post, error) {
	rows, err := db.Query("SELECT guestName, content FROM guestbook ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Name, &post.Content); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
