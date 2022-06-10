package configserver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//mysql接口
var MySqlMaster *gorm.DB

//连接字符串
var MysqlDsn string

func Gorm() {

	sql, err := gorm.Open("mysql", MysqlDsn)
	MySqlMaster = sql
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//自动判断是否存在表 并创建表
	e := MySqlMaster.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Config{})
	if e != nil {
		fmt.Println(e)
	}
}
