package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8765"
	}

	host := address + ":" + port

	g := gin.Default()

	var lastPostRequest *http.Request
	var lastPostBody string

	g.GET("/request", func(c *gin.Context) {
		text, err := json.Marshal(lastPostRequest)
		if err != nil {
			c.String(500, err.Error())
		} else {
			c.String(200, string(text))
		}
	})

	g.GET("/body", func(c *gin.Context) {
		c.String(200, lastPostBody)
	})

	g.POST("/*path", func(c *gin.Context) {
		lastPostRequest = c.Request
		lastPostBody, _ = getBodyAsString(c.Request.Body)
		c.String(200, "ok")
	})

	fmt.Println("listening on", host)

	g.Run(host)
}

func getBodyAsString(body io.ReadCloser) (string, error) {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	out := string(b)
	return out, nil
}
