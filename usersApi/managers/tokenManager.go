package managers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
	"usersApi/models"
)

func GenerateToken(user models.User) (error, string) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"iss":    "usersApi",
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
		"nbf":    time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error("Error while generating token", err)
		return err, ""
	}
	return nil, tokenString
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Error(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return false
		}
	}
	return token.Valid
}
