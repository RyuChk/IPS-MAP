package models

import "time"

type Building struct {
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	OriginLat   float64   `bson:"origin_lat"`
	OriginLong  float64   `bson:"origin_long"`
	FloorList   []Floor   `bson:"-"`
	CreatedAt   time.Time `bson:"created_at"` //Create time
	UpdatedAt   time.Time `bson:"updated_at"` //Update time
}

type Floor struct {
	Name        string    `bson:"name"`        //Map Name
	Description string    `bson:"description"` //Floor Description
	Floor       int       `bson:"floor"`       //Floor number
	Symbol      string    `bson:"symbol"`
	Building    string    `bson:"building"`   //Name of the building
	IsAdmin     bool      `bson:"is_admin"`   //Only admin can view
	CreatedAt   time.Time `bson:"created_at"` //Create time
	UpdatedAt   time.Time `bson:"updated_at"` //Update time
}

type FloorDetail struct {
	Key       string    `bson:"key"` //Key
	Info      Floor     `bson:"-"`
	RoomList  []Room    `bson:"room"`
	CreatedAt time.Time `bson:"created_at"` //Create time
	UpdatedAt time.Time `bson:"updated_at"` //Update time
}

type Room struct {
	RoomID      string    `bson:"room_id"`     // Room ID
	Name        string    `bson:"name"`        // Name of room
	Description string    `bson:"description"` // Description
	Latitude    float64   `bson:"latitude"`    // lat location
	Longitude   float64   `bson:"longitude"`   // long location
	CreatedAt   time.Time `bson:"created_at"`  //Create time
	UpdatedAt   time.Time `bson:"updated_at"`  //Update time
}
