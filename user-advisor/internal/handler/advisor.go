package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-advisor/internal/middleware"
	"user-advisor/internal/model"
	"user-advisor/internal/model/req"
	"user-advisor/internal/model/resp"
	"user-advisor/internal/service"
)

func AdvisorRegister(c *gin.Context) {
	var request req.AdvisorRegister
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	err := service.AdvisorRegister(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "success"})
}

func AdvisorLogin(c *gin.Context) {
	var request req.AdvisorLogin
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	user, err := service.AdvisorLogin(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	j := middleware.NewJWT()
	claims := j.CreateClaims(middleware.BaseClaims{
		Phone:    user.Phone,
		Password: user.Password,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": errors.New("获取token失败").Error()})
		return
	}
	response := resp.AdvisorLogin{
		Phone: user.Phone,
		Name:  user.Name,
		Token: token,
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

func AdvisorUpgrade(c *gin.Context) {
	var request req.AdvisorUpgrade
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	advisor := &model.Advisor{
		Name:     request.Name,
		Bio:      request.Bio,
		Gender:   request.Gender,
		Birth:    request.Birth,
		Phone:    claim.Phone,
		Password: claim.Password,
		Coin:     request.Coin,
	}
	err := service.AdvisorUpgrade(advisor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "upgrade successfully"})
}

func UpgradeWorkStatus(c *gin.Context) {
	var request req.UpgradeWorkStatus
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	var advisor model.Advisor
	advisor.WorkStatus = request.WorkStatus
	advisor.Phone = claim.Phone
	advisor.Password = claim.Password
	err := service.UpgradeWorkStatus(advisor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "upgrade successfully"})
}

func UpgradeReading(c *gin.Context) {
	var request req.UpgradeReading
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	var reading model.Reading
	reading.AdvisorPhone = claim.Phone
	reading.Details = request.Details
	reading.UserCost = request.UserCost
	err := service.UpgradeReading(reading)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "upgrade successfully"})
}

func GetAllOrderByAdvisor(c *gin.Context) {
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	orders, err := model.GetAllOrderByAdvisor(claim.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	if orders == nil {
		c.JSON(http.StatusOK, map[string]interface{}{"data": "has no order form"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": orders})
}

func DealOrder(c *gin.Context) {
	var request req.DealOrder
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	order, err := model.GetOneOrder(claim.Phone, request.UserPhone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	if order.Status == 1 {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": errors.New("the order has been deal").Error()})
		return
	}
	err = service.DealOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "deal the order form successfully"})
}

func AdvisorGetCoinStream(c *gin.Context) {
	var request req.AdvisorGetCoinStream
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	coinStream := service.AdvisorGetCoinStream(claim.Phone, request.Order)
	c.JSON(http.StatusOK, map[string]interface{}{"data": coinStream})
}
