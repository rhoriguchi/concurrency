package entity

import "time"

type User struct {
	Name      string
	Birthdate time.Time
	Email     string
}
