package models

import (
	"github.com/jinzhu/gorm"
)

type SystemState int

const (
	Unconfigured SystemState = iota + 1
	Configured
)

type Master struct {
	gorm.Model

	Name     string
	Ip       string
	User     string
	Password string
}

type Configuration struct {
	Name string
}

type System struct {
	gorm.Model

	State    SystemState
	Master   Master
	MasterID int

	Config   Configuration
	ConfigID int
}
