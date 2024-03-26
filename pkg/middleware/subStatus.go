package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyPremiumStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("Session")
		if err != nil {
			ctx.AbortWithStatusJSON(401, err)
			return
		}

		token, _, err := new(jwt.Parser).ParseUnverified(cookie, jwt.MapClaims{})
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {

			expirationDate, err := time.Parse("2006-01-02 15:04:05", claims["subStatus"].(string))
			if err != nil {
				ctx.JSON(500, "Error parsing expiration date from token claims:"+err.Error())
			}

			if time.Now().Before(expirationDate) {
				ctx.Next()
			} else {
				ctx.AbortWithStatusJSON(401, err)
				return
			}

		}
	}
}
