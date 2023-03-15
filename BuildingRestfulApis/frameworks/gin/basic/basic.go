package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	/*Get takes a route and a handler func
	handler takes the gin context obj*/
	r.GET("/pingTime", func(c *gin.Context) { //context holds the info of the individual request
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run(":2007") // listen and serve

}
