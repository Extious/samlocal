package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-advisor/internal/model"
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
