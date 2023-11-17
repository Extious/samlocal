package model

import (
	"go-web/global"
)

const TableUser = "User1"

type User struct {
	Name     string `dynamo:"name" json:"name"`
	Phone    string `dynamo:"phone,hash" json:"phone"`
	Password string `dynamo:"password" json:"password"`
}

func TableUserCreate() error {
	return global.DB.CreateTable(TableUser, &User{}).Run()
}
