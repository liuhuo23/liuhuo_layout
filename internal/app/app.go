package app

import (
	"fmt"
	"os"

	config "{{cookiecutter.project_name}}/internal/conf"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type App struct {
	Config *config.AppConfig
	Echo   *echo.Echo
	DB     *gorm.DB
}

func NewApp(config *config.AppConfig, echo *echo.Echo, db *gorm.DB) *App {
	return &App{
		Config: config,
		Echo:   echo,
		DB:     db,
	}
}

func (a *App) Run() {
	fmt.Println(fmt.Sprintf(
		"Starting server 0.0.0.0:%d",
		a.Config.Port,
	))
}

func NewAppConfig() *config.AppConfig {
	// 读取配置文件
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 解析YAML到AppConfig
	var cfg config.AppConfig
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	//  读取密钥文件
	secretData, err := os.ReadFile("config/secrets.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(secretData, &cfg); err != nil {
		panic(err)
	}
	// 为避免复制锁值，传递 cfg 的指针
	fmt.Println(&cfg)
	return &cfg
}

func NewDatabaseConfig(config *config.AppConfig) *config.DatabaseConfig {
	return config.Database
}
