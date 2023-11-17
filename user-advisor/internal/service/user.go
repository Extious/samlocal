package service

import (
	"errors"
	"user-advisor/global"
	"user-advisor/internal/model"
	"user-advisor/internal/model/req"
	"user-advisor/internal/model/resp"
)

func UserRegister(req *req.UserRegister) error {
	user1 := model.UserScanByPhone(req.Phone)
	if user1 != nil {
		return errors.New("the phone has been registered")
	}
	var user = &model.User{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
	}
	err := user.Insert()
	if err != nil {
		return err
	}
	return nil
}

func UserLogin(req *req.UserLogin, ip string) (*model.User, error) {
	user1 := model.UserScanByPhone(req.Phone)
	if user1 == nil {
		return nil, errors.New("the phone has not been registered")
	}
	for _, v := range user1.Ip {
		if v == ip {
			if user1.Password != req.Password {
				return nil, errors.New("the password is error")
			}
			return user1, nil
		}
	}
	if user1.Password != req.Password {
		return nil, errors.New("the password is error")
	}
	user1.Ip = append(user1.Ip, ip)
	user1.Insert()
	return user1, nil
}

func UserUpgrade(user *model.User) error {
	err := user.Insert()
	if err != nil {
		return err
	}
	return nil
}

func GetAdvisorList() []resp.GetAdvisorList {
	advisors := model.GetAllAdvisor()
	if advisors == nil {
		return nil
	}
	var advisorsInfo []resp.GetAdvisorList
	for _, k := range advisors {
		advisorsInfo = append(advisorsInfo, resp.GetAdvisorList{
			Name:  k.Name,
			Phone: k.Phone,
			Bio:   k.Bio,
		})
	}
	return advisorsInfo
}

func GetOneAdvisor(phone string) (resp.GetOneAdvisor, error) {
	advisor := model.AdvisorScanByPhone(phone)
	if advisor == nil {
		return resp.GetOneAdvisor{}, errors.New("not found the advisor")
	}
	advisorInfo := resp.GetOneAdvisor{
		Name:        advisor.Name,
		Phone:       advisor.Phone,
		WorkStatus:  advisor.WorkStatus,
		Bio:         advisor.Bio,
		Gender:      advisor.Gender,
		Birth:       advisor.Birth,
		OrderNumber: advisor.OrderNumber,
		Experience:  advisor.Experience,
		Star:        advisor.Star,
		CommentNum:  advisor.CommentNum,
	}
	return advisorInfo, nil
}

func BookOrder(order model.OrderForm) error {
	reading, err := model.GetOneReading(order.AdvisorPhone, order.Details)
	if err != nil {
		return err
	}
	user := model.UserScanByPhone(order.UserPhone)
	if user.Coin-reading.UserCost >= 0 {
		user.Coin = user.Coin - reading.UserCost
		user.CoinStream = append(user.CoinStream, -reading.UserCost)
		user.Insert()
	} else {
		return errors.New("the coin is not enough")
	}
	advisor := model.AdvisorScanByPhone(order.AdvisorPhone)
	advisor.OrderNumber++
	advisor.Insert()
	order.Insert()
	return nil
}

func DeliverComment(comment model.CommentInfo) error {
	reading, err := model.GetOneReading(comment.AdvisorPhone, comment.Details)
	if err != nil {
		return errors.New("no such reading")
	}
	reading.Comment = append(reading.Comment, comment)
	reading.Insert()
	comment.Insert()
	advisor := model.AdvisorScanByPhone(comment.AdvisorPhone)
	advisor.CommentNum++
	advisor.Star = (advisor.Star*float64(advisor.CommentNum-1) + comment.Star) / float64(advisor.CommentNum)
	advisor.Insert()
	return nil
}

func GiveReward(userPhone, advisorPhone string, coin float64) error {
	user := model.UserScanByPhone(userPhone)
	if user.Coin-coin >= 0 {
		user.Coin = user.Coin - coin
		user.CoinStream = append(user.CoinStream, -coin)
		user.Insert()
	} else {
		return errors.New("the coin is not enough")
	}
	advisor := model.AdvisorScanByPhone(advisorPhone)
	advisor.Coin = advisor.Coin + coin
	advisor.CoinStream = append(advisor.CoinStream, +coin)
	advisor.Insert()
	return nil
}

func StarAdvisor(userPhone, advisorPhone string) error {
	user := model.UserScanByPhone(userPhone)
	user.StarAdvisors = append(user.StarAdvisors, advisorPhone)
	user.Insert()
	return nil
}

func GetStaredAdvisor(userPhone string) []string {
	user := model.UserScanByPhone(userPhone)
	return user.StarAdvisors
}

func UserGetCoinStream(phone string, order bool) []float64 {
	user := model.UserScanByPhone(phone)
	if order == true {
		return user.CoinStream
	} else {
		return global.ReverseSlice(user.CoinStream)
	}
}
