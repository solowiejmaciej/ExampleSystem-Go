package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"usersApi/initializers"
)

func TestRabbitMQ(context *gin.Context) {
	log.Info("TestRabbitMQ")
}

func TestRedis(context *gin.Context) {
	initializers.RedisClient.Set("key", 123, 5555550)
	val, err := initializers.RedisClient.Get("key").Result()
	if err != nil {
		log.Error(err)
	}
	log.Info(val)
}
