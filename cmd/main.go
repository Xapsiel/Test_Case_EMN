package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"mobileTest_Case/internal/handler"
	"mobileTest_Case/internal/models"
	"mobileTest_Case/internal/repository"
	"mobileTest_Case/internal/service"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Warn("Ошибка инициализации переменных конфигурации")
	}
	if err := godotenv.Load(); err != nil {
		logrus.Warn("Ошибка инициализации переменных окружения")
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Warn(err.Error())
	}
	repos := repository.NewRepository(db)
	logrus.Info("Определение переменной бд")
	services := service.NewService(repos)
	logrus.Info("Определение переменной сервисов")
	handlers := handler.NewHandler(services)
	logrus.Info("Определение маршрутов")
	srv := new(models.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Warn(err.Error())
	}
	logrus.Info("Запущен сервер на порту :" + viper.GetString("port"))

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
