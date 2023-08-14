package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	"github.com/jutionck/interview-bootcamp-apps/utils/security"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	RequireToken(userRole ...string) gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwtSecurity security.JwtSecurity
}

func (a *authMiddleware) RequireToken(userRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			common.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			common.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		claims, err := a.jwtSecurity.VerifyAccessToken(tokenString)
		if err != nil {
			common.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		c.Set("claims", claims)

		validRole := false
		for _, role := range userRole {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			common.SendErrorResponse(c, http.StatusForbidden, "Forbidden Resource")
			return
		}

		c.Next()
	}
}

func (a *authMiddleware) RefreshToken() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func NewAuthMiddleware(jwtSecurity security.JwtSecurity) AuthMiddleware {
	return &authMiddleware{jwtSecurity: jwtSecurity}
}
