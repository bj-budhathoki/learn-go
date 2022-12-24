package infrastructure

import "os"

// Env has environment stored
type Env struct {
	SERVER_PORT string
	ENVIRONMENT string
	DB_HOST     string
	DB_DRIVER   string
	API_SECRET  string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

func (env *Env) LoadEnv() {
	env.SERVER_PORT = os.Getenv("SERVER_PORT")
	env.ENVIRONMENT = os.Getenv("ENVIRONMENT")

	env.DB_USER = os.Getenv("DB_USER")
	env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	env.DB_HOST = os.Getenv("DB_HOST")
	env.DB_PORT = os.Getenv("DB_PORT")
	env.DB_NAME = os.Getenv("DB_NAME")
}
