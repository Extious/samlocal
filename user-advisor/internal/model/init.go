package model

import (
	"user-advisor/global"
)

func TableCheck() {
	list, err := global.DB.ListTables().All()
	if err != nil {
		panic(err)
	}
	hash := make(map[string]bool)
	for _, v := range list {
		hash[v] = true
	}

	if _, ok := hash[TableUser]; !ok {
		err = TableUserCreate()
		if err != nil {
			panic(err)
		}
	}

	if _, ok := hash[TableAdviser]; !ok {
		err = TableAdviserCreate()
		if err != nil {
			panic(err)
		}
	}

	if _, ok := hash[TableReading]; !ok {
		err = TableReadingCreate()
		if err != nil {
			panic(err)
		}
	}

	if _, ok := hash[TableOrderForm]; !ok {
		err = TableOrderFormCreate()
		if err != nil {
			panic(err)
		}
	}

	if _, ok := hash[TableComment]; !ok {
		err = TableCommentCreate()
		if err != nil {
			panic(err)
		}
	}
}
