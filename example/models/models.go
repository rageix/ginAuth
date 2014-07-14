package models

type Users struct {
	Id       int
	Email    string `orm:"size(255)"`
	Password string `orm:"size(60)"`
}
