package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env Env) Database {
	dsn := fmt.Sprintf("host=postgresdb_container user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB, env.POSTGRES_PORT, env.SSL_MODE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.POSTGRES_DB + ";")
	if err != nil {
		log.Println("Url: ", dsn)
		log.Panic(err)
	}
	log.Println("Database connection established")
	return Database{
		DB: db,
	}
}
