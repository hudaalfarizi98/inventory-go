package main

import (
	"fmt"

	"catalog-go/modules/users"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r = gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./scripts", false)))
	r.LoadHTMLGlob("templates/*")

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	users.Initial(r)

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/ping", func(c *gin.Context) {
			fmt.Println("PING TEST")
		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			fmt.Println("PING ANALYTICS")
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("TEST")
	}
}
