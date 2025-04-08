package handler

import (
	"fmt"
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))


	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "OK",
		})
	})
	
	server.Handle(w, r)
}