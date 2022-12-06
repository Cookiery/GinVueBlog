package middleware

import (
	"errors"
	"main/commond/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JWTKey = []byte("123")

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

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
func VerifyToken(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if key, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		return key, nil
	} else {
		return nil, nil
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

		key, err := VerifyToken(token[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    statusCode,
				"message": "token 出错",
			})
			c.Abort()
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
