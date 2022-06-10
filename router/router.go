package router

import (
	"ConfigServer/configserver"
	"github.com/gin-gonic/gin"
)

func InitRouter(R *gin.Engine) {
	R.GET("/set/:user/:pwd/:ConfigName/:row", configserver.SetConfig)
	R.GET("/get/:user/:pwd/:ConfigName", configserver.GetConfig)
}
