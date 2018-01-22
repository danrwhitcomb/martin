package models

import (
	"github.com/jinzhu/gorm"
)

type SystemState int

const (
	Available SystemState = iota + 1
	Owned
)

type Master struct {
	gorm.Model

	name     string
	ip       string
	user     string
	password string
}

type Configuration struct {
	name string
}

type System struct {
	gorm.Model

	state  SystemState
	master Master
	config Configuration
}
