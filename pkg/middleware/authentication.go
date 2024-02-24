package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/pkg/firebaseConnection"
)

// func AuthenticateFirebaseIdToken() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authClient := firebaseConnection.CreateFirebaseClient()

// 		headerIdToken := ctx.GetHeader("Bearer")
// 		_, err := authClient.VerifyIDToken(ctx, headerIdToken)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(401, err)
// 			return
// 		} else {
// 			ctx.Next()
// 		}
// 	}
// }

func VerifySessionCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("session")
		if err != nil {
			ctx.AbortWithStatusJSON(401, err)
			return
		}
		authClient := firebaseConnection.CreateFirebaseClient()
		_, err = authClient.VerifySessionCookieAndCheckRevoked(ctx, cookie)
		if err != nil {
			ctx.AbortWithStatusJSON(498, err)
			return
		} else {
			ctx.Next()
		}
	}
}
