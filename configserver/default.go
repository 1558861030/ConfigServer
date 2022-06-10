package configserver

import (
	"github.com/jinzhu/gorm"
)

type Config struct {
	gorm.Model
	User       string `gorm:"index"`
	Pwd        string `gorm:"index"`
	ConfigName string `gorm:"index"`
	Row        string
}
