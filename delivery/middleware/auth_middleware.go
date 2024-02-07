package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ismailash/be-enigma-laundry/config"
	"github.com/ismailash/be-enigma-laundry/utils/common"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService common.JwtToken
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func NewAuthMiddleware(jwtService common.JwtToken) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authHeader authHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		tokenString := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set(config.UserSession, claims["userId"])

		// cek validitas role
		var validRole bool
		for _, role := range roles {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Resource"})
			return
		}

		ctx.Next()
	}
}
