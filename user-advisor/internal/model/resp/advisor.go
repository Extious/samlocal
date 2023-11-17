package resp

type (
	AdvisorLogin struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Token string `json:"token"`
	}
)
