package main

import (
	"headphonista/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Router(r)

	r.Run(":80")
}
