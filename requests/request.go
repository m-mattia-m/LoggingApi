package requests

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewRequest(c *gin.Context) {
	status := c.PostForm("status")
	url := c.PostForm("url")
	application := c.PostForm("application")
	title := c.PostForm("title")
	message := c.PostForm("message")
	success := c.GetBool(c.PostForm("success"))
	currentRequest := addRequest(status, url, application, title, message, success)
	c.JSON(200, currentRequest)
}

func addRequest(status string, url string, application string, title string, message string, success bool) Request {
	var funcRequest = new(Request)
	funcRequest.Id = uuid.New().String()
	funcRequest.Status = status
	funcRequest.Url = url
	funcRequest.Application = application
	funcRequest.Title = title
	funcRequest.Message = message
	funcRequest.RequestTime = time.Now()
	funcRequest.Success = success
	fmt.Println("------------------------------------")
	fmt.Println("New Request successfully created:")
	fmt.Println("------------------------------------")
	fmt.Println(funcRequest)
	fmt.Println("\tId\t\t", funcRequest.Id, "\n\tStatus\t", funcRequest.Status, "\n\tUrl\t", funcRequest.Url, "\n\tApplication\t", funcRequest.Application, "\n\tTitle\t\t", funcRequest.Title, "\n\tMessage\t\t", funcRequest.Message, "\n\tRequestTime\t", funcRequest.RequestTime, "\n\tSuccess\t", funcRequest.Success)
	fmt.Println("------------------------------------")

	return *funcRequest
}

type Request struct {
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Url         string    `json:"url"`
	Application string    `json:"application"`
	Title       string    `json:"title"`
	Message     string    `json:"message"`
	RequestTime time.Time `json:"time"`
	Success     bool      `json:"success"`
}
