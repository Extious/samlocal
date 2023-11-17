package model

import "go-web/global"

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
}
