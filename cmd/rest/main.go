package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"github.com/jakskal/koperasi-v2/config"
	"github.com/jakskal/koperasi-v2/server"
	_ "github.com/joho/godotenv/autoload"

	"gorm.io/gorm"
)

func main() {

	conf := config.Get()

	dbHost := conf.DBHost
	dbPort := conf.DBPort
	dbUser := conf.DBUser
	dbName := conf.DBName
	dbPassword := conf.DBPassword
	dsn := fmt.Sprintf("host=%s user=%s database=%s port=%v password=%s sslmode=disable", dbHost, dbUser, dbName, dbPort, dbPassword)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Print("failed to connect to database", err)
	} else {
		log.Print("Connection to database established")
	}
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()
	if err := server.StartServer(conf, db); err != nil {
		log.Fatal("failed to start server", err)

	}
}

func closeDB(db *gorm.DB) {
	dbInstance, _ := db.DB()
	dbInstance.Close()
}
