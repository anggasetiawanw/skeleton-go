package entity

import "time"

type Devices struct {
	ID         string
	Imei       string
	Address    string
	Status     bool
	Type       string
	Created_At time.Time
	Updated_At time.Time
}