package main

import (
	"context"
	"fmt"
	"go_api/internal/comment/db"
)

func Run() error {
	fmt.Println("Application running")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Fail to connect database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return nil
	}
	return nil
}

func main() {
	fmt.Println("Hello world")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
