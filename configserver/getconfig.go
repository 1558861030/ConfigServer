package configserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConfig(c *gin.Context) {
	config := Config{}
	config.User = c.Param("user")
	config.Pwd = c.Param("pwd")
	config.ConfigName = c.Param("ConfigName")

	MySqlMaster.Where("User = ? AND Pwd = ? AND Config_Name = ? ", config.User, config.Pwd, config.ConfigName).First(&config)

	c.String(http.StatusOK, config.Row)
}
