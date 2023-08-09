package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var p = ":8888"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/", gin.Dir(".", true))

	if len(os.Args) > 1 {
		p = fmt.Sprintf(":%s", os.Args[1])
	}

	fmt.Println("Listening and serving HTTP on", p)
	r.Run(p)
}
