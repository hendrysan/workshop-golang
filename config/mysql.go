package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db_user := viper.GetString("Database.User")
	db_password := viper.GetString("Database.Password")
	db_host := viper.GetString("Database.Host")
	db_port := viper.GetInt("Database.Port")
	db_name := viper.GetString("Database.Name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)

	mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return mysql, nil
}
