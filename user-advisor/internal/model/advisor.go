package model

import (
	"time"
	"user-advisor/global"
)

type WorkStatus int8

const (
	Open  WorkStatus = 0
	Close WorkStatus = 1
)

const TableAdviser = "Advisor"

type Advisor struct {
	Name        string     `json:"name" dynamo:"name"`
	Gender      Gender     `json:"gender" dynamo:"gender"`
	Phone       string     `json:"phone" dynamo:"phone,hash"`
	Password    string     `json:"password" dynamo:"password"`
	Birth       string     `json:"birth" dynamo:"birth"`
	Bio         string     `json:"bio" dynamo:"bio"`
	Coin        float64    `json:"coin" dynamo:"coin"`
	Status      Status     `json:"status" dynamo:"status"`
	WorkStatus  WorkStatus `json:"work_status" dynamo:"work_status"`
	OrderNumber int64      `json:"order_number" dynamo:"order_number"`
	Star        float64    `json:"star" dynamo:"star"`
	CommentNum  int64      `json:"comment_num" dynamo:"comment_num"`
	Experience  int32      `json:"experience" dynamo:"experience"`
	CreateTime  time.Time  `json:"create_time" dynamo:"create_time"`
	CoinStream  []float64  `dynamo:"coin_stream" json:"coin_stream"`
}

func TableAdviserCreate() error {
	return global.DB.CreateTable(TableAdviser, &Advisor{}).Run()
}

func (advisor *Advisor) Insert() error {
	table := global.DB.Table(TableAdviser)
	return table.Put(advisor).Run()
}

func DeleteAdviser(name string) error {
	table := global.DB.Table(TableAdviser)
	return table.Delete("name", name).Run()
}

func AdvisorScanByPhone(phone string) *Advisor {
	table := global.DB.Table(TableAdviser)
	var advisors []Advisor
	table.Scan().All(&advisors)
	for _, v := range advisors {
		if v.Phone == phone {
			return &v
		}
	}
	return nil
}

func GetAllAdvisor() []Advisor {
	table := global.DB.Table(TableAdviser)
	var advisors []Advisor
	table.Scan().All(&advisors)
	return advisors
}
