package models

import (
	"time"
)

type OnlineUser struct {
	DisplayName string    `json:"display_name"`
	Building    string    `json:"building"`
	Floor       int       `json:"floor"`
	Coordinate  Position  `json:"coordinate"`
	Timestamp   time.Time `json:"timestamp"`
}

type Position struct {
	X float64 `json:"x"` // X position in meter
	Y float64 `json:"y"` // Y position in meter
	Z float64 `json:"z"` // Z position in floor
}
