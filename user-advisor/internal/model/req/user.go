package req

import (
	"user-advisor/internal/model"
)

type (
	UserRegister struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	UserLogin struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	UserUpgrade struct {
		Name   string       `json:"name"`
		Gender model.Gender `json:"gender"`
		Birth  string       `json:"birth"`
		Bio    string       `json:"bio"`
		Coin   float64      `json:"coin"`
	}
	GetOneAdvisor struct {
		Phone string `json:"phone"`
	}
	BookOrder struct {
		AdvisorPhone string `json:"advisor_phone"`
		Details      string `json:"details"`
	}
	GiveReward struct {
		AdvisorPhone string  `json:"advisor_phone"`
		Coin         float64 `json:"coin"`
	}
	StarAdvisor struct {
		AdvisorPhone string `json:"advisor_phone"`
	}
	UserGetCoinStream struct {
		Order bool `json:"order"`
	}
)
