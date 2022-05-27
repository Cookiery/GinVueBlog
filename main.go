package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

type Server struct {
	AppMode  string
	HTTPPort string
}

func main() {
	var app map[string]Server
	if _, err := toml.DecodeFile("./conf/app.toml", &app); err != nil {
		fmt.Println(app)
	}
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello"})
	})
	r.Run()
}
