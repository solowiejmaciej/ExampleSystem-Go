package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"usersApi/initializers"
)

func CachePage(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.RequestURI()
		val, err := initializers.RedisClient.Get(key).Result()
		if err == nil {
			log.Info("Cache hit, returning the cached data.")
			var jsonData interface{}
			json.Unmarshal([]byte(val), &jsonData)
			c.JSON(http.StatusOK, jsonData)
			c.Abort()
			return
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		body := blw.body.String()
		err = initializers.RedisClient.Set(key, body, duration).Err()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
