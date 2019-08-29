package model

type User struct {
	Id int64
	Name string `xorm:"varchar(25) notnull unique"`
}
