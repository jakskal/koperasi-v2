package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jakskal/koperasi-v2/pkg/token"
)

const xClientId = "X-ClientId"

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

const UserIDClaim = "UserID"
const roleClaim = "RoleID"

// AuthWithRoleRequired scenes software that allows request to communicate and interact with application by authentication.
func AuthWithRoleRequired(permittedRoles ...float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Token not found",
			})
			return
		} else if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token, need bearer token",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)

		userID := claims[UserIDClaim]
		role := claims[roleClaim]

		// Set example variable
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid token, user id not exist",
			})
		}

		if role == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid token, role not exist",
			})
			return
		}

		if !isRoleHasRight(role.(float64), permittedRoles...) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden request",
			})
			return
		}

		c.Set(UserIDClaim, int(userID.(float64)))
		c.Set(roleClaim, role)

		// Parse JWT Claims
		var claimInfo token.Claims
		_, _, err = new(jwt.Parser).ParseUnverified(tokenString, &claimInfo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "token parsed invalid",
				"err":     err.Error(),
			})
			return
		}

		ctx := context.WithValue(c.Request.Context(), TokenInfoContextKey, claimInfo)
		c.Request = c.Request.WithContext(ctx)
		// before request
		c.Next()

	}
}

func AuthAdminAcess() gin.HandlerFunc {
	return AuthWithRoleRequired(0, 1, 2)
}

func AuthAcess() gin.HandlerFunc {
	return AuthWithRoleRequired(0, 1, 2, 3)
}

func isRoleHasRight(role float64, roles ...float64) bool {
	isHasRight := false
	for _, value := range roles {
		if role == value {
			isHasRight = true
		}
	}
	return isHasRight

}
