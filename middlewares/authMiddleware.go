package middlewares

import (
	"fmt"
	"net/http"
	helpers "restaurant/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// -so for authenication with JWT we need to check if the token is in the request header

		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			ctx.Abort()
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.First_name)
		ctx.Set("last_name", claims.Last_name)
		ctx.Set("uid", claims.Uid)

		ctx.Next()
		//-if its in the header we validate the token
		//-and after validating we pick the data out of the token and pass it to the route

	}
}
