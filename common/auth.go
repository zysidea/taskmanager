package common

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"log"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
)

//使用非对称加密
const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

//privateKey 用来注册JWT，publicKey 用来验证HTTP请求
var (
	verifyKey *rsa.PublicKey
	signKey *rsa.PrivateKey
)

type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

//获取JWT
func GenerateJWT(name, role string) (string, error) {
	claims := AppClaims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}



	//用rsa256
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//验证JWT的中间件
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {


	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		//JWT验证错误
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			//JWT过期了
			case jwt.ValidationErrorExpired:
				DisplayAppError(
					w,
					err,
					"Access Token is expired,get a new Token",
					http.StatusUnauthorized,
				)
				return

			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Access Token",
					http.StatusInternalServerError,
				)
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing the Access Token",
				http.StatusInternalServerError,
			)
		}
	}

	if token.Valid {
		context.Set(r, "user", token.Claims.(*AppClaims).UserName)
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
