package main

import (
	"fmt"
	"liuhuo23/liuos/internal/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 读取配置文件
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %v", err))
	}

	// 获取数据库配置
	dsn := viper.GetString("data.database.source")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败: %v", err))
	}

	// 自动迁移模型
	err = db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.UserRole{},
		&model.RolePermission{},
	)
	if err != nil {
		panic(fmt.Errorf("数据库迁移失败: %v", err))
	}

	fmt.Println("数据库迁移完成")
}
