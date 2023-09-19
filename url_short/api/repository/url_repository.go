package repository

import "github.com/bj-budhathoki/learn-go/url_short/infrastructure"

type UrlRepository struct {
	database infrastructure.Database
}

func NewUrlRepository(database infrastructure.Database) UrlRepository {
	return UrlRepository{
		database: database,
	}
}
