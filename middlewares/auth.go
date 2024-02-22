package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ichbinnichts/events-api/utils"
)

func Authenticate(context *gin.Context) {

	// Get token from header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}

	context.Set("userId", userId)
	context.Next()
}
