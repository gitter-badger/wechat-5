package models

type Admin struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	Name     string `xorm:"VARCHAR(45)"`
	Passwd   string `xorm:"VARCHAR(45)"`
	Email    string `xorm:"VARCHAR(45)"`
	Mobile   string `xorm:"VARCHAR(11)"`
	RoleType int    `xorm:"INT(11)"`
}
