package main

import (
	"headphonista/src/bootstrap"
	"headphonista/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Router(r)

	r.Run(":80")
	defer bootstrap.CloseDB()
}
