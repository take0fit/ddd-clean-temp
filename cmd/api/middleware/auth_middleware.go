package middleware

import (
	"crypto/rsa"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func AuthMiddleware(publicKey *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// "Bearer "を除去してトークン文字列を取得
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// トークンの署名検証を行い、ペイロードを取得
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		// `sub`クレームからユーザーIDを取得してリクエストに添付
		if userId, ok := claims["sub"].(string); ok {
			c.Set("userID", userId)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user ID claim is missing"})
			return
		}

		c.Next() // 次のハンドラへ処理を継続
	}
}
