package models

import (
	"fmt"
	"time"
)

type Comment struct {
	Id         string    `xorm:"not null pk VARCHAR(255)"`
	Comment    string    `xorm:"not null VARCHAR(200)"`
	CreateTime time.Time `xorm:"DATETIME created"`
	PostId     string    `xorm:"not null VARCHAR(255)"`
	UpdateTime time.Time `xorm:"DATETIME updated"`
	UserId     string    `xorm:"not null VARCHAR(255)"`
}

func (c *Comment)BeforeInsert()  {
	fmt.Println("before insert",*c)
}
func (c *Comment) AfterInsert() {
	fmt.Println("after insert",*c)
}