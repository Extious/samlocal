package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user-advisor/internal/middleware"
	"user-advisor/internal/model"
	"user-advisor/internal/model/req"
	"user-advisor/internal/model/resp"
	"user-advisor/internal/service"
)

func UserRegister(c *gin.Context) {
	var request req.UserRegister
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	err := service.UserRegister(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "success"})
}

func UserLogin(c *gin.Context) {
	var request req.UserLogin
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	user, err := service.UserLogin(&request, c.ClientIP())
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
	response := resp.UserLogin{
		Phone: user.Phone,
		Name:  user.Name,
		Token: token,
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

func UserUpgrade(c *gin.Context) {
	var request req.UserUpgrade
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	user := &model.User{
		Name:     request.Name,
		Bio:      request.Bio,
		Gender:   request.Gender,
		Birth:    request.Birth,
		Phone:    claim.Phone,
		Password: claim.Password,
		Coin:     request.Coin,
	}
	err := service.UserUpgrade(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "upgrade successfully"})
}

func GetAdvisorList(c *gin.Context) {
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	advisors := service.GetAdvisorList()
	if advisors == nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": errors.New("not found the advisors").Error()})
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": advisors})
}

func GetOneAdvisor(c *gin.Context) {
	var request req.GetOneAdvisor
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	advisor, err := service.GetOneAdvisor(request.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": advisor})
}

func BookOrder(c *gin.Context) {
	var request req.BookOrder
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	var orderForm model.OrderForm
	orderForm.Details = request.Details
	orderForm.AdvisorPhone = request.AdvisorPhone
	orderForm.UserPhone = claim.Phone
	orderForm.Status = model.OrderToDo
	err := service.BookOrder(orderForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "book the order successfully"})
}

func GetAllOrderByUser(c *gin.Context) {
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	orders, err := model.GetAllOrderByUser(claim.Phone)
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

func DeliverComment(c *gin.Context) {
	var request req.DeliverComment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	var comment = model.CommentInfo{
		AdvisorPhone:   request.AdvisorPhone,
		Details:        request.Details,
		UserPhone:      claim.Phone,
		Star:           request.Star,
		CommentTime:    time.Now(),
		CommentContent: request.CommentContent,
	}
	err := service.DeliverComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "comment successfully"})
}

func GiveReward(c *gin.Context) {
	var request req.GiveReward
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	err := service.GiveReward(claim.Phone, request.AdvisorPhone, request.Coin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "give the reward successfully"})
}

func StarAdvisor(c *gin.Context) {
	var request req.StarAdvisor
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	err := service.StarAdvisor(claim.Phone, request.AdvisorPhone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "star successfully"})
}

func GetStaredAdvisor(c *gin.Context) {
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	advisors := service.GetStaredAdvisor(claim.Phone)
	if advisors == nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": errors.New("no stared advisor").Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": advisors})
}

func UserGetCoinStream(c *gin.Context) {
	var request req.UserGetCoinStream
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"data": "the request is error"})
		return
	}
	claim := middleware.GetUserInfo(c)
	if claim == nil {
		c.JSON(http.StatusProxyAuthRequired, map[string]interface{}{"data": errors.New("auth wrong").Error()})
		return
	}
	coinStream := service.UserGetCoinStream(claim.Phone, request.Order)
	c.JSON(http.StatusOK, map[string]interface{}{"data": coinStream})
}
