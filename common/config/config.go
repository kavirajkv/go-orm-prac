package config

import (
	"github.com/spf13/viper"
)


var (
	DB_USERNAME=""
	DB_PASSWORD=""
	DB_PORT=""
	DB_HOST=""
	DB_NAME=""
	PORT=""
)


func InitConfig() {
	viper.SetConfigName(".env") 
	viper.SetConfigType("env")   
	viper.AddConfigPath(".")     
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv() // Automatically read environment variables

	DB_USERNAME = viper.GetString("DB_USERNAME")
	DB_PASSWORD = viper.GetString("DB_PASSWORD")
	DB_PORT = viper.GetString("DB_PORT")
	DB_HOST = viper.GetString("DB_HOST")
	DB_NAME = viper.GetString("DB_NAME")
	PORT = viper.GetString("PORT")

}
