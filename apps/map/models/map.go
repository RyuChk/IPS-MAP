package models

import "time"

type MapURL struct {
	Name      string    `bson:"name"`
	Floor     int       `bson:"floor"`
	ImageURL  string    `bson:"image_url"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
