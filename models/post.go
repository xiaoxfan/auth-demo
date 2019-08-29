package models

import (
	"time"
)

type Post struct {
	Id           string    `xorm:"not null pk VARCHAR(255)"`
	Category     int       `xorm:"INT(11)"`
	Content      string    `xorm:"VARCHAR(2000)"`
	CreateTime   time.Time `xorm:"DATETIME"`
	Title        string    `xorm:"not null VARCHAR(255)"`
	UpdateTime   time.Time `xorm:"DATETIME"`
	UserId       string    `xorm:"not null VARCHAR(255)"`
	CommentCount int       `xorm:"INT(11)"`
}
