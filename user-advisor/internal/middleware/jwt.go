package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user-advisor/global"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type JWT struct {
	LoginKey []byte
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	Phone    string
	Password string
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.JwtLoginKey),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {

	now := time.Now().Unix()

	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: 86400, // buffer time 1 day buffer time will get a new token refresh token. In this case, a user will have two valid tokens, but only one will be left at the front end and the other will be lost.
		StandardClaims: jwt.StandardClaims{
			NotBefore: now - 1000,    // effective time of signature
			ExpiresAt: now + 604800,  // Expiration time 7 days profile
			Issuer:    global.Issuer, // the publisher of the signature
			IssuedAt:  now,           //issue time
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.LoginKey)
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// We have jwt authentication header information to return token information when Authorization logs in. Here,
		//the front end needs to store the token in cookie or local localStorage,
		//but you need to negotiate the expiration time with the back end.
		//You can agree to refresh the token or log in again.
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("未登录或非法访问").Error()})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("授权已过期").Error()})
				c.Abort()
				return
			}
			c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.LoginKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := c.Request.Header.Get("Authorization")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("Failed to obtain parsing information from jwt from Context of Gin. Please check whether Authorization exists in the request header and whether claims is the specified structure.")
	}
	return claims, err
}

func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
