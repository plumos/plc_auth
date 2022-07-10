package myhttp

import (
	"github.com/gin-gonic/gin"
)

func HttpStart() {
	r := gin.Default()
	api := r.Group("api")
	api.POST("/dataEncrypt", dataEncrypt)
	api.POST("/dataDecrypt", dataDecrypt)
	api.POST("/sign", sign)
	api.POST("/verifySign", verifySign)
	api.POST("/resultNotify", resultNotify)

}
