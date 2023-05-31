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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", env.POSTGRES_HOST, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB, env.POSTGRES_PORT, env.SSL_MODE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Url: ", dsn)
		log.Panic(err)
	}
	// stmt := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", env.POSTGRES_DB)
	// if rs := db.Exec(stmt); rs.Error != nil {
	// 	log.Println(rs.Error)
	// }
	log.Println("Database connection established")
	return Database{
		DB: db,
	}
}
