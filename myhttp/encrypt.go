package myhttp

import (
	"github.com/gin-gonic/gin"
	"plc_auth/model"
)

func dataEncrypt(c *gin.Context) {
	var req model.EncryptReq
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

}

func dataDecrypt(c *gin.Context) {

}
