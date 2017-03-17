package common

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
)

//使用非对称加密
const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

//privateKey 用来注册JWT，publicKey 用来验证HTTP请求
var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}

//获取JWT
func GenerateJWT(name, role string) (string, error) {
	//用rsa256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	//设置过期时间
	t.Claims["expired"] = time.Now().Add(20 * time.Minute).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//验证JWT的中间件
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	header := r.Header
	tokenString := header["Authorization"][0]
	//Bearer token....，截取token
	tokenString = strings.Split(tokenString, " ")[1]
	//对token进行验证
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
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
