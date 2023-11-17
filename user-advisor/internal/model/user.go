package model

import (
	"time"
	"user-advisor/global"
)

type Gender int8

const (
	Man   Gender = 0
	Woman Gender = 1
)

type Status int8

const (
	Normal Status = 0
	Banned Status = 1
)

const TableUser = "User"

type User struct {
	Name         string    `dynamo:"name" json:"name"`
	Gender       Gender    `dynamo:"gender" json:"gender"`
	Phone        string    `dynamo:"phone,hash" json:"phone"`
	Password     string    `dynamo:"password" json:"password"`
	Birth        string    `dynamo:"birth" json:"birth"`
	Bio          string    `dynamo:"bio" json:"bio"`
	Coin         float64   `dynamo:"coin" json:"coin"`
	Status       Status    `dynamo:"status" json:"status"`
	StarAdvisors []string  `dynamo:"star_advisors" json:"star_advisors"`
	CreateTime   time.Time `dynamo:"create_time" json:"create_time"`
	CoinStream   []float64 `dynamo:"coin_stream" json:"coin_stream"`
	Ip           []string  `dynamo:"ip" json:"ip"`
}

func TableUserCreate() error {
	return global.DB.CreateTable(TableUser, &User{}).Run()
}

func (user *User) Insert() error {
	table := global.DB.Table(TableUser)
	return table.Put(user).Run()
}

func DeleteUser(name string) error {
	table := global.DB.Table(TableUser)
	return table.Delete("name", name).Run()
}

func UserScanByPhone(phone string) *User {
	table := global.DB.Table(TableUser)
	var users []User
	table.Scan().All(&users)
	for _, v := range users {
		if v.Phone == phone {
			return &v
		}
	}
	return nil
}
