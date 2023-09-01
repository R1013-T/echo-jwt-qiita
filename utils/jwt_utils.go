package utils

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"time"
)

type JwtCustomClaims struct {
	id uint
	jwt.RegisteredClaims
}

var signingKey = []byte(Getenv("JWT_SECRET"))

var JwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(JwtCustomClaims)
	},
	SigningKey:  signingKey,
	TokenLookup: "cookie:token",
}

func GenerateToken(id uint) string {
	claims := &JwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(signingKey)
	if err != nil {
		return ""
	}

	return t
}

func GetClaims(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	return claims
}
