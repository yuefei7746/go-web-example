package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/yuefei7746/go-web-example/utils"
)

var (
	AuthorizationAccount = gin.Accounts{
		"yuefei7746": "123123",
	}
)

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString(utils.StringToBytes(base))
}
