package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/takt-corp/gin-logrus"
	"time"
	"usersApi/controllers"
	"usersApi/initializers"
	"usersApi/middleware"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConfigureLogs()
	initializers.ConnectToDB()
	initializers.Migrate()
	initializers.ConnectToRabbitMq()
	initializers.ConnectToRedis()
}
func main() {
	r := gin.New()
	r.Use(ginlogrus.LoggerMiddleware(ginlogrus.LoggerMiddlewareParams{}))
	r.GET("/api/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.POST("/api/user", controllers.AddUser)
	//r.GET("/api/user/:userId", middleware.RequireToken, middleware.CachePage(time.Hour), controllers.GetUserById)
	r.GET("/api/user/:userId", middleware.CachePage(time.Hour), controllers.GetUserById)
	r.POST("/api/token", controllers.GenerateToken)

	r.GET("/api/rabbit", controllers.TestRabbitMQ)
	r.GET("/api/redis", controllers.TestRedis)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
