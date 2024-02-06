package models

import "time"

type Map struct {
	Name        string    `bson:"name"`        //Map Name
	Description string    `bson:"description"` //Floor Description
	Number      int       `bson:"number"`      //Floor number
	Symbol      string    `bson:"symbol"`
	Building    string    `bson:"building"`   //Name of the building
	IsAdmin     bool      `bson:"is_admin"`   //Only admin can view
	CreatedAt   time.Time `bson:"created_at"` //Create time
	UpdatedAt   time.Time `bson:"updated_at"` //Update time
}

type MapImageURL struct {
	Key       string    `bson:"key"`        //<Building>-<Symbol>
	MapDetail Map       `bson:"-"`          //Map detail
	URL       string    `bson:"url"`        //Map image URL
	CreatedAt time.Time `bson:"created_at"` //Create time
	UpdatedAt time.Time `bson:"updated_at"` //Update time
}
