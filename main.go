package main

import (
	"github.com/bj-budhathoki/learn-go/bootstrap"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
