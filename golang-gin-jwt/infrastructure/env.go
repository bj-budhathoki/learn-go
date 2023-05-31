package infrastructure

import (
	"fmt"
	"os"
)

type Env struct {
	POSTGRES_HOST            string
	POSTGRES_PORT            string
	POSTGRES_USER            string
	POSTGRES_PASSWORD        string
	POSTGRES_DB              string
	DATABASE_URL             string
	CLIENT_ORIGIN            string
	PGADMIN_DEFAULT_EMAIL    string
	PGADMIN_DEFAULT_PASSWORD string
	JWT_SECRET               string
	JWT_EXPIRED_IN           string
	JWT_MAXAGE               string
	SSL_MODE                 string
	Environment              string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	fmt.Println("load env")
	env.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	env.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	env.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	env.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	env.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	env.CLIENT_ORIGIN = os.Getenv("CLIENT_ORIGIN")
	env.PGADMIN_DEFAULT_EMAIL = os.Getenv("PGADMIN_DEFAULT_EMAIL")
	env.PGADMIN_DEFAULT_PASSWORD = os.Getenv("PGADMIN_DEFAULT_PASSWORD")
	env.JWT_SECRET = os.Getenv("JWT_SECRET")
	env.JWT_EXPIRED_IN = os.Getenv("JWT_EXPIRED_IN")
	env.JWT_MAXAGE = os.Getenv("JWT_MAXAGE")
	env.SSL_MODE = os.Getenv("SSL_MODE")
	env.Environment = os.Getenv("Environment")

}
