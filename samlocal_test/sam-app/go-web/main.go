package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"go-web/global"
	"go-web/internal/handler"
	"go-web/internal/model"
	"net/http"
)

var ginLambda *ginadapter.GinLambda

// Handler main proxy func
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	global.InitDynamoDB()
	model.TableCheck()
	if ginLambda == nil {
		r := gin.Default()
		api := r.Group("/api")

		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{"data": "pong"})
		})
		api.POST("/db_put", handler.DbTest)
		api.GET("/getData", handler.DbGet)
		ginLambda = ginadapter.New(r)
	}
	return ginLambda.ProxyWithContext(ctx, req)
}
func main() {
	lambda.Start(Handler)
}
