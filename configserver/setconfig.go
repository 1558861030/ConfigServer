package configserver

import (
	Log "github.com/1558861030/LogToLogServer/LogToLogServer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetConfig(c *gin.Context) {
	config := Config{}

	config.User = c.Param("user")
	config.Pwd = c.Param("pwd")
	config.ConfigName = c.Param("ConfigName")
	config.Row = c.Param("row")

	if config.User == "" && config.Pwd == "" && config.ConfigName == "" && config.Row == "" {
		c.String(400, "参数不能为空", config)
	}

	configs := Config{}
	MySqlMaster.Where("User = ? AND Pwd = ? AND Config_Name = ? ", config.User, config.Pwd, config.ConfigName).First(&configs)
	//判断数据库中是否存在不存在则创建 否则则更新数据
	if configs.ID == 0 {
		resp := MySqlMaster.Create(&config)
		if resp.Error != nil {
			Log.Error(resp.Error)
		}
	} else {
		resp := MySqlMaster.Model(&configs).Update("Row", config.Row)
		if resp.Error != nil {
			Log.Error(resp.Error)
		}
	}

	c.String(http.StatusOK, "ok")
}
