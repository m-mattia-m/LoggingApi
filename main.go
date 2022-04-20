package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var users []User

func main() {

	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	r.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now(),
		})
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "on",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":      "on",
			"time":        time.Now(),
			"uptime":      time.Since(time.Now()),
			"Gin-Version": gin.Version,
			"Gin-Mode":    gin.Mode(),
			"Gin-Debug":   gin.IsDebugging(),
		})
	})

	r.POST("/registration", func(c *gin.Context) {
		postFirstname := c.PostForm("firstname")
		postLastname := c.PostForm("lastname")
		postEmail := c.PostForm("email")
		postUsername := c.PostForm("username")
		postUrl := c.PostForm("url")
		postApplication := c.PostForm("application")

		if postFirstname == "" || postLastname == "" || postEmail == "" || postUsername == "" || postUrl == "" || postApplication == "" {
			c.JSON(400, "Bitte alle Felder ausfüllen")
		} else {
			var postUser User = newUser(postUsername, postFirstname, postLastname, postEmail, postUrl, postApplication)
			users = append(users, postUser)
			c.JSON(200, gin.H{
				"yourToken": postUser.Token,
				"firstname": postUser.Firstname,
				"lastname":  postUser.Lastname,
				"email":     postUser.Email,
				"username":  postUser.Username,
			})
		}
	})

	authorized.GET("/getusers", func(c *gin.Context) {
		usersArr := make([]User, len(users))
		for i := 0; i < len(users); i++ {
			usersArr[i] = users[i]
		}
		c.JSON(200, users)
	})

	authorized.POST("/logging", func(c *gin.Context) {
		httpstatus := c.PostForm("httpstatus")
		url := c.PostForm("url")
		application := c.PostForm("application")
		loggingTitle := c.PostForm("title")
		message := c.PostForm("message")
		senderToken := c.PostForm("sender")

		fmt.Println("sendertoken: " + senderToken)

		currentUser, _ := addRequestToUser(senderToken, newRequest(httpstatus, url, application, loggingTitle, message, true))
		c.JSON(200, currentUser.Request[len(currentUser.Request)-1])
	})

	authorized.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "nothing to see here"})
	})

	r.Run(":3000") // for a hard coded port

}

// zu verbessern: Performantere Suche in den Slices (users) -> evtl. mit Byte-Arrays (noch genauer ansehen)
func addRequestToUser(userToken string, currentRequest Request) (User, error) {
	for i := 0; i < len(users); i++ {
		fmt.Println("user[i].Token: " + users[i].Token)
		fmt.Println("current userToken: " + userToken)
		if users[i].Token == userToken {
			fmt.Println("3. Token is equal")
			users[i].Request = append(users[i].Request, currentRequest)
			return users[i], nil
		}
	}
	return User{}, nil
}

func newRequest(status string, url string, application string, title string, message string, success bool) Request {
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

func newUser(username string, firstname string, lastname string, email string, url string, application string) User {
	var funcUser = new(User)
	funcUser.Id = uuid.New().String()
	funcUser.Username = username
	funcUser.Firstname = firstname
	funcUser.Lastname = lastname
	funcUser.Email = email
	funcUser.Token = uuid.New().String()
	funcUser.Request = append(funcUser.Request, newRequest("200", url, application, "Registration", "User '"+funcUser.Id+"' wurde erfolgreich erstellt", true))
	fmt.Println("------------------------------------")
	fmt.Println("New User successfully created:")
	fmt.Println("------------------------------------")
	// fmt.Println(funcUser)
	fmt.Println("\tId\t\t", funcUser.Id, "\n\tUsername\t", funcUser.Username, "\n\tFirstname\t", funcUser.Firstname, "\n\tLastname\t", funcUser.Lastname, "\n\tEmail\t\t", funcUser.Email, "\n\tToken\t\t", funcUser.Token, "\n\tRequests\t", funcUser.Request)
	fmt.Println("------------------------------------")

	return *funcUser
}

type User struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	Request   []Request `json:"Request"`
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
