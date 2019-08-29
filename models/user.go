package models

import (
	"time"
)

type User struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	CreateTime time.Time `xorm:"DATETIME"`
	Email      string    `xorm:"VARCHAR(64)"`
	Mobile     string    `xorm:"unique VARCHAR(13)"`
	Password   string    `xorm:"VARCHAR(255)"`
	Role       int       `xorm:"INT(11)"`
	Salt       string    `xorm:"VARCHAR(255)"`
	UpdateTime time.Time `xorm:"DATETIME"`
	Username   string    `xorm:"not null unique VARCHAR(64)"`
}
