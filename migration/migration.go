package migration

import (
	"fmt"
	"go-fiber-gorm/config"
	"go-fiber-gorm/model"
	"log"
)

func RunMigration() {

	err := config.DB.AutoMigrate(
		&model.User{},
		&model.Transaction{},
	)

	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Database Migrated")
}
