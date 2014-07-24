package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

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

	stack := &RequestStack{}

	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		var out string
		for i, req := range *stack {
			out += strconv.Itoa(i) + ": " + req.Details.URL.Path + "\n"
		}
		c.String(200, out)
	})

	g.POST("/*path", func(c *gin.Context) {
		body, _ := getBodyAsString(c.Request.Body)
		r := &Request{
			Details: c.Request,
			Body:    body,
		}
		stack.add(r)
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
