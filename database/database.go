package database

import (
	"fmt"
	"log"
	"os"

	"filtering-sample/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct{
	DB *gorm.DB
}

var DB DBInstance

func ConnectDB(){
	data:= fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432  sslmode=disable TimeZone=Asia/Shanghai",os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	
	db , err :=gorm.Open(postgres.Open(data), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	} )

	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")
	db.Logger= db.Logger.LogMode(logger.Info)

	log.Println("Migrating the schema")
	db.AutoMigrate(&models.Listing{}, &models.Location{}, &models.Utility{}, &models.Features{})
	DB = DBInstance{
		DB: db, 
	}
}
