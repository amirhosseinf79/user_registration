package database

import (
	"fmt"
	"log"
	"time"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormconnection(connStr string, debug bool) *gorm.DB {
	gormConfig := &gorm.Config{}
	if debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	var db *gorm.DB
	var err error

	for {
		fmt.Println("Connectiong to SQL-DB...")
		db, err = gorm.Open(postgres.Open(connStr), gormConfig)
		if err != nil {
			fmt.Println("failed to connect database:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	return db
}
