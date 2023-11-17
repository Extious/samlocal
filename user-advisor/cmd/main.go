package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-advisor/global"
	"user-advisor/internal/handler"
	"user-advisor/internal/middleware"
	"user-advisor/internal/model"
)

func main() {
	global.InitDynamodb()
	model.TableCheck()
	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"data": "pong"})
	})
	api.POST("/db_put", handler.DbTest)
	userLogin := api.Group("/user")
	{
		userLogin.POST("/register", handler.UserRegister)
		userLogin.POST("/login", handler.UserLogin)
	}
	advisorLogin := api.Group("/advisor")
	{
		advisorLogin.POST("/register", handler.AdvisorRegister)
		advisorLogin.POST("/login", handler.AdvisorLogin)
	}
	user := api.Group("/user")
	user.Use(middleware.JWTAuth())
	{
		user.GET("/get_advisor_list", handler.GetAdvisorList)
		user.POST("/get_one_advisor", handler.GetOneAdvisor)
		user.POST("/upgrade", handler.UserUpgrade)
		user.POST("/book_order", handler.BookOrder)
		user.GET("/get_all_orders", handler.GetAllOrderByUser)
		user.POST("/deliver_comment", handler.DeliverComment)
		user.POST("/give_a_reward", handler.GiveReward)
		user.POST("/star_advisor", handler.StarAdvisor)
		user.GET("/get_stared_advisor", handler.GetStaredAdvisor)
		user.POST("/get_coin_stream", handler.UserGetCoinStream)
	}
	advisor := api.Group("advisor")
	advisor.Use(middleware.JWTAuth())
	{
		advisor.POST("/upgrade", handler.AdvisorUpgrade)
		advisor.POST("/work_status", handler.UpgradeWorkStatus)
		advisor.POST("/upgrade_reading", handler.UpgradeReading)
		advisor.GET("/get_all_orders", handler.GetAllOrderByAdvisor)
		advisor.POST("/deal_order", handler.DealOrder)
		advisor.POST("/get_coin_stream", handler.AdvisorGetCoinStream)
	}
	r.Run(":3000")
}
