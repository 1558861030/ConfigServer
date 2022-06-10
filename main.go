package main

import (
	"ConfigServer/configserver"
	"ConfigServer/router"
	"github.com/1558861030/LogToLogServer/LogToLogServer"
	"github.com/gin-gonic/gin"
)

func main() {
	//配置日志服务器ip
	LogToLogServer.LogServer.IpPost = "127.0.0.1:9100"
	//是否记录每个用户返回的信息
	LogToLogServer.LogServer.ResplogBTN = true
	//项目名称
	LogToLogServer.LogServer.ProjectName = "mydwl"
	//应用名称
	LogToLogServer.LogServer.App = "configserver"

	//配置configserver mysql 连接
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	configserver.MysqlDsn = "configserver:password@tcp(127.0.0.1:3306)/configserver?charset=utf8mb4&parseTime=True&loc=Local"
	configserver.Gorm()

	//新建gin路由
	Gin := gin.New()
	//注册日志服务器
	LogToLogServer.InitLogToLogServer(Gin)
	//注册路由
	router.InitRouter(Gin)
	Gin.Run(":9101")

}
