package model

import (
	"fmt"
	"github.com/guregu/dynamo"
	"time"
	"user-advisor/global"
)

type (
	CommentInfo struct {
		AdvisorPhone   string    `json:"advisor_phone" dynamo:"advisor_phone,hash"`
		Details        string    `json:"details" dynamo:"details"`
		UserPhone      string    `json:"user_phone" dynamo:"user_phone"`
		Star           float64   `json:"star" dynamo:"star"`
		CommentTime    time.Time `json:"comment_time" dynamo:"comment_time,range"`
		CommentContent string    `json:"comment_content" dynamo:"comment_content"`
	}
	Reading struct {
		AdvisorPhone string        `json:"advisor_phone" dynamo:"advisor_phone,hash"`
		Details      string        `json:"details" dynamo:"details,range"`
		UserCost     float64       `json:"user_cost" dynamo:"user_cost"`
		Comment      []CommentInfo `json:"comment" dynamo:"comment"`
	}
	OrderForm struct {
		AdvisorPhone string `json:"advisor_phone" dynamo:"advisor_phone,hash"`
		UserPhone    string `json:"user_phone" dynamo:"user_phone,range"`
		Details      string `json:"details" dynamo:"details"`
		Status       int8   `json:"status" dynamo:"status"`
	}
)

const TableComment = "Comment"
const TableReading = "Reading"
const TableOrderForm = "OrderForm"

const OrderDone = 1
const OrderToDo = 0

func TableReadingCreate() error {
	err := global.DB.CreateTable(TableReading, &Reading{}).Run()
	fmt.Println(err)
	return err
}

func TableOrderFormCreate() error {
	return global.DB.CreateTable(TableOrderForm, &OrderForm{}).Run()
}

func TableCommentCreate() error {
	err := global.DB.CreateTable(TableComment, &CommentInfo{}).Index(dynamo.Index{
		Name:           "CommentTime",
		HashKey:        "advisor_phone",
		HashKeyType:    dynamo.StringType,
		RangeKey:       "details",
		RangeKeyType:   dynamo.StringType,
		Local:          true,
		ProjectionType: dynamo.AllProjection,
	}).Run()
	if err != nil {
		return err
	}
	return nil
}

func (reading *Reading) Insert() error {
	table := global.DB.Table(TableReading)
	return table.Put(reading).Run()
}

func (orderForm *OrderForm) Insert() error {
	table := global.DB.Table(TableOrderForm)
	return table.Put(orderForm).Run()
}

func (comment *CommentInfo) Insert() error {
	table := global.DB.Table(TableComment)
	return table.Put(comment).Run()
}

func GetOneReading(advisorPhone, details string) (*Reading, error) {
	table := global.DB.Table(TableReading)
	var reading Reading
	err := table.Get("advisor_phone", advisorPhone).Range("details", dynamo.Equal, details).One(&reading)
	if err != nil {
		return nil, err
	}
	return &reading, nil
}

func GetOneOrder(advisorPhone, userPhone string) (*OrderForm, error) {
	table := global.DB.Table(TableOrderForm)
	var order OrderForm
	err := table.Get("advisor_phone", advisorPhone).Range("user_phone", dynamo.Equal, userPhone).One(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func GetAllOrderByAdvisor(phone string) ([]OrderForm, error) {
	table := global.DB.Table(TableOrderForm)
	var orders []OrderForm
	err := table.Get("advisor_phone", phone).All(&orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetAllOrderByUser(phone string) ([]OrderForm, error) {
	table := global.DB.Table(TableOrderForm)
	var orders []OrderForm
	var order []OrderForm
	err := table.Scan().All(&orders)
	if err != nil {
		return nil, err
	}
	for _, v := range orders {
		if v.UserPhone == phone {
			order = append(order, v)
		}
	}
	return order, nil
}
