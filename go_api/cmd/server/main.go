package main

import (
	"fmt"
	"go_api/internal/db"
)

func Run() error {
	fmt.Println("Application running")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect ot the  database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
