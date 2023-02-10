package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var MY_APPLICATION_NAME = "RESTAPI ECOMMERCE"
var SIGNED_KEY = []byte("LZV7XdXJQ8yfroozepLh9fEm1S0NCZ")
var EXP_DATE = &jwt.NumericDate{Time: time.Now().Add(24 * time.Hour)}

type MYClaims struct {
	*jwt.RegisteredClaims
	ID       int    `json:"id"`
	Username string `json:"username" form:"username"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" form:"username"`
}

func JWTMiddleware(c *gin.Context) {

	getToken := c.Request.Header.Get("Authorization")
	if getToken == "" {
		msg := exception.ToErrorMsg("Failed, please login first. Your token is empty.", exception.Unauthorized)
		exception.ErrorHandler(c, msg)
		return
	} else {
		authorizerHeader := strings.SplitAfter(getToken, "Bearer ")
		if len(authorizerHeader) == 1 {
			msg := exception.ToErrorMsg("Your token is empty.", exception.BadRequest)
			exception.ErrorHandler(c, msg)
			return
		}

		token, err := jwt.Parse(authorizerHeader[1], func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}
			c.Header("Authorization", fmt.Sprintf("Bearer %s", token.Raw))
			c.SetCookie("Authorization", fmt.Sprintf("Bearer %s", token.Raw), 10, "/", "localhost:8080", true, true)

			return SIGNED_KEY, nil
		})
		if err != nil {
			msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
			exception.ErrorHandler(c, msg)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !token.Valid || !ok {
			msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
			exception.ErrorHandler(c, msg)
			return
		}

		return
	}
}

func generateJWTToken(user *User) (string, error) {
	var claims = &MYClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    MY_APPLICATION_NAME,
			ExpiresAt: EXP_DATE,
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		ID:       user.ID,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(SIGNED_KEY)
	if err != nil {
		return "", err
	}

	resultToken, err := json.Marshal(signedString)
	if err != nil {
		return "", errors.New("Cannot Marshal")
	}

	return string(resultToken), nil
}

func NewGenerateJWTToken(user *User) (string, error) {
	return generateJWTToken(user)
}
