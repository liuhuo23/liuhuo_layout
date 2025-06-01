package data

import (
	"fmt"
	config "liuhuo23/liuos/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb(data *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		data.Username,
		data.Password,
		data.Host, data.Port, data.DbName,
	)
	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return sqlDB
}
