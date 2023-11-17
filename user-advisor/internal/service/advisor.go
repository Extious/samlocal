package service

import (
	"errors"
	"user-advisor/global"
	"user-advisor/internal/model"
	"user-advisor/internal/model/req"
)

func AdvisorRegister(req *req.AdvisorRegister) error {
	advisor1 := model.AdvisorScanByPhone(req.Phone)
	if advisor1 != nil {
		return errors.New("the phone has been registered")
	}
	var advisor = &model.Advisor{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
	}
	err := advisor.Insert()
	if err != nil {
		return err
	}
	return nil
}

func AdvisorLogin(req *req.AdvisorLogin) (*model.Advisor, error) {
	advisor1 := model.AdvisorScanByPhone(req.Phone)
	if advisor1 == nil {
		return nil, errors.New("the phone has not been registered")
	}
	if advisor1.Password != req.Password {
		return nil, errors.New("the password is error")
	}
	return advisor1, nil
}

func AdvisorUpgrade(advisor *model.Advisor) error {
	err := advisor.Insert()
	if err != nil {
		return err
	}
	return nil
}

func UpgradeWorkStatus(advisor model.Advisor) error {
	err := advisor.Insert()
	if err != nil {
		return err
	}
	return nil
}

func UpgradeReading(reading model.Reading) error {
	err := reading.Insert()
	if err != nil {
		return err
	}
	return nil
}

func DealOrder(order *model.OrderForm) error {
	order.Status = 1
	order.Insert()
	reading, err := model.GetOneReading(order.AdvisorPhone, order.Details)
	if err != nil {
		return err
	}
	advisor := model.AdvisorScanByPhone(order.AdvisorPhone)
	advisor.Coin = advisor.Coin + reading.UserCost
	advisor.CoinStream = append(advisor.CoinStream, +reading.UserCost)
	advisor.Experience++
	advisor.Insert()
	return nil
}

func AdvisorGetCoinStream(phone string, order bool) []float64 {
	advisor := model.AdvisorScanByPhone(phone)
	if order == true {
		return advisor.CoinStream
	} else {
		return global.ReverseSlice(advisor.CoinStream)
	}
}
