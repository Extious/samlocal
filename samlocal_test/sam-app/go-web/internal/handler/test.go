package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/internal/model"
	"net/http"
)

func DbTest(c *gin.Context) {
	test := &model.DbTest{}
	err := c.BindJSON(test)
	fmt.Println(test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	err = test.DbTest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": "success"})
	return
}

func DbGet(c *gin.Context) {
	test := []*model.DbTest{}
	test, err := model.DbGet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"data": test})
	return
}
