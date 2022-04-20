package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	profile "github.com/lvornholt/go-profiles"
	"strings"
	"time"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	Email    string
	UserType string
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(email string, userType string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("unable to parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}
	return
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}
		jwtWrapper := JwtWrapper{
			SecretKey: profile.GetStringValue("jwtConfig.secretKey"),
			Issuer:    profile.GetStringValue("jwtConfig.issuer"),
		}
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Next()
	}
}
