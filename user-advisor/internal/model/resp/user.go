package resp

import "user-advisor/internal/model"

type (
	UserLogin struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Token string `json:"token"`
	}
	GetAdvisorList struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Bio   string `json:"bio"`
	}
	GetOneAdvisor struct {
		Name        string           `json:"name"`
		Phone       string           `json:"phone"`
		Bio         string           `json:"bio"`
		Gender      model.Gender     `json:"gender"`
		Birth       string           `json:"birth"`
		WorkStatus  model.WorkStatus `json:"work_status"`
		OrderNumber int64            `json:"order_number"`
		Star        float64          `json:"star"`
		CommentNum  int64            `json:"comment_num"`
		Experience  int32            `json:"experience"`
	}
)
