package common

import (
	"net/http"
	"io/ioutil"
	"github.com/labstack/gommon/log"
)

//使用非对称加密
const (
	privateKeyPath ="keys/app.rsa"
	publicKeyPath="keys/app.rsa.pub"
)
//privateKey 用来注册JWT，publicKey 用来验证HTTP请求
var(
	verifyKey,signKey []byte
)
func initKeys(){
	var err error
	signKey,err=ioutil.ReadFile(privateKeyPath)
	if err!=nil {
		log.Fatalf("[initKeys]: %s\n",err)
	}
	verifyKey,err=ioutil.ReadFile(publicKeyPath)
	if err!=nil {
		log.Fatalf("[initKeys]: %s\n",err)
		panic(err)
	}
}
func Authorize(w http.ResponseWriter,r *http.Request){

}
