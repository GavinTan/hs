package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var port = "8888"

func main() {
	r := gin.Default()
	r.StaticFS("/", gin.Dir(".", true))

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	r.Run(fmt.Sprintf(":%s", port))
}
