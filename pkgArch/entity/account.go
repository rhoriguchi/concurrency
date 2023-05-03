package entity

import "sync"

type Account struct {
	Balance float64
	Owner   *User
	Mu      sync.RWMutex
}
