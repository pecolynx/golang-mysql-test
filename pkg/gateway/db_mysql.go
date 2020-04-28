package gateway

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func OpenMySQL(config *MySQLConfig) (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Asia%%2FTokyo", config.User, config.Password, config.Host, config.Port, config.Database))
}
