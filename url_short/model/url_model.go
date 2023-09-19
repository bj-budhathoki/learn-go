package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type URLBody struct {
	LongURL string `json:"log_url"`
}
type Url struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UrlCode   string             `bson:"url_code"`
	LongURL   string             `bson:"long_url"`
	ShortURL  string             `bson:"short_url"`
	CreatedAt int64              `bson:"created_at"`
	ExpiredAt int64              `bson:"expired_at"`
}
