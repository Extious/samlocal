package req

import "user-advisor/internal/model"

type (
	AdvisorRegister struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	AdvisorLogin struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	AdvisorUpgrade struct {
		Name   string       `json:"name"`
		Gender model.Gender `json:"gender"`
		Birth  string       `json:"birth"`
		Bio    string       `json:"bio"`
		Coin   float64      `json:"coin"`
	}
	UpgradeWorkStatus struct {
		WorkStatus model.WorkStatus `json:"work_status"`
	}
	UpgradeReading struct {
		Details  string  `json:"details"`
		UserCost float64 `json:"user_cost"`
	}
	DealOrder struct {
		UserPhone string `json:"user_phone"`
	}
	DeliverComment struct {
		AdvisorPhone   string  `json:"advisor_phone"`
		Details        string  `json:"details"`
		Star           float64 `json:"star"`
		CommentContent string  `json:"comment_content"`
	}
	AdvisorGetCoinStream struct {
		Order bool `json:"order"`
	}
)
