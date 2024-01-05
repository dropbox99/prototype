package config

import (
	"fmt"
	"log"
	"os"
	"prototype/lib/env"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigMysql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var (
	cfg *ConfigMysql
)

func init() {
	cfg = &ConfigMysql{
		Host:     env.String("Database.Host", ""),
		Port:     env.String("Database.Port", ""),
		User:     env.String("Database.User", ""),
		Password: env.String("Database.Pass", ""),
		Database: env.String("Database.Name", ""),
	}
}

func NewMysql() (*gorm.DB, error) {
	// init connection mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	log.Printf("INFO: Connected to DB")

	return db, nil

}
