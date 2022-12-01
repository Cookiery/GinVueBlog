package middleware

import (
	"main/commond/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JWTKey = []byte("123")

type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// CreateToken create token by JWTKey
func CreateToken(userName string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	claims := Claims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GinVueBlog", // 发行人
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(JWTKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// VerifyToken verify token
func VerifyToken(token string) (*Claims, int) {
	jwtToken, _ := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if key, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR
	}
}

// JWTToken jwt controler
func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		statusCode := errmsg.SUCCSE
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			statusCode = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    statusCode,
				"message": errmsg.ErrMsg(statusCode),
			})
			c.Abort()
			return
		}

		token := strings.SplitN(tokenHeader, " ", 2)
		if len(token) != 2 && token[0] != "Bearer" {
			statusCode = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    statusCode,
				"message": errmsg.ErrMsg(statusCode),
			})
			c.Abort()
			return
		}

		key, statusCode := VerifyToken(token[1])
		if statusCode == errmsg.ERROR {
			statusCode = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    statusCode,
				"message": errmsg.ErrMsg(statusCode),
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt.Unix() {
			statusCode = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    statusCode,
				"message": errmsg.ErrMsg(statusCode),
			})
			c.Abort()
			return
		}

		c.Set("userName", key.UserName)
		c.Next()
	}
}
