package main

import (
	"bookspreadLogging/users"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(handler)
	// Define the router
	r := gin.Default()

	// Define the user-routes
	r.POST("/registration", users.Registration)
	r.GET("/getUsers", users.BasicAuth, users.GetUsers)
	r.GET("/deleteUser/:id", users.BasicAuth, users.GetUser)
	r.GET("/deleteUser", users.BasicAuth, func(c *gin.Context) {
		c.JSON(400, "Send an ID of a user with. Example: /deleteUser/id")
	})
	r.GET("/getUser/:id", users.BasicAuth, users.GetUser)
	r.GET("/getUser", users.BasicAuth, func(c *gin.Context) {
		c.JSON(400, "Send an ID of a user with. Example: /getUser/id")
	})
	r.POST("/editUser/:id", users.BasicAuth, users.EditUser)
	r.POST("/editUser", users.BasicAuth, func(c *gin.Context) {
		c.JSON(400, "Send an ID of a user with. Example: /getUser/id")
	})

	// Define the request-routes
	r.POST("/request", users.BasicAuth, users.CreateRequst)

	r.Run(":3000")
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "text/plain"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              "Hello, World!",
		IsBase64Encoded:   false,
	}, nil
}
