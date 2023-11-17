package model

import (
	"fmt"
	"user-advisor/global"
)

type DbTest struct {
	ID   int    `dynamo:"id" json:"id"`
	Name string `dynamo:"name" json:"name"`
}

func (person *DbTest) DbTest() error {
	table := global.DB.Table("test")
	fmt.Println("1111111")
	err := table.Put(person).Run()
	if err != nil {
		fmt.Println("err:")
		fmt.Println(err)
		return err
	}
	return nil
}
