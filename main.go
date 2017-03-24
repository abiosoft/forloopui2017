package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abiosoft/river"
)

func main() {
	rv := river.New().Renderer(river.JSONRender)
	rv.Handle("/", river.NewEndpoint().Get("/", hello))

	{
		e := river.NewEndpoint()
		e.Get("/", handleListPosts)
		e.Post("/", handleNewPost)
		rv.Handle("/guestbook", e)
	}
	log.Fatal(rv.Run(":" + os.Getenv("PORT")))
}

func hello(c *river.Context) {
	c.Render(200, "Hello ForLoop")
}

func handleListPosts(c *river.Context) {
	posts, err := getPosts()
	if err != nil {
		c.Render(http.StatusInternalServerError, err.Error())
		return
	}
	c.Render(200, posts)
}

func handleNewPost(c *river.Context) {
	var post Post
	if err := c.DecodeJSONBody(&post); err != nil {
		c.Render(http.StatusBadRequest, "Bad Request")
	}
	if err := createPost(post); err != nil {
		c.Render(http.StatusInternalServerError, err.Error())
		return
	}
	c.Render(201, post)
}
