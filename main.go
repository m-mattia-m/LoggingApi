package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	var users []User

	r := gin.Default()
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

	r.GET("/data", func(c *gin.Context) {
		a := make([]string, 2)
		a[0] = "John"
		a[1] = "Sam"
		j, err := json.Marshal(a)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(j))
		}
	})

	r.GET("/getusers", func(c *gin.Context) {

		usersArr := make([]User, len(users))

		for i := 0; i < len(users); i++ {
			usersArr[i] = users[i]
		}

		// userJson, err := json.Marshal(usersArr)
		// input := []byte(userJson)
		// if err != nil {
		// 	c.JSON(400, gin.H{"error-name": "userArray is empty (nil)", "error-message": err.Error()})
		// } else {
		// 	buf := &bytes.Buffer{}
		// 	if err := json.Indent(buf, input, "", "\t"); err != nil {
		// 		c.JSON(400, gin.H{"error": err.Error()})
		// 	}
		// 	// userJson2, err := json.MarshalIndent(userJson, "", "\t")
		// 	// if err != nil {
		// 	// 	c.JSON(200, string(userJson2))
		// 	// 	fmt.Println(string(userJson2))

		// 	// } else {
		// 	// 	c.JSON(200, string(userJson2))
		// 	// 	fmt.Println(string(userJson2))
		// 	// }

		// 	c.JSON(200, buf.String())
		// 	fmt.Println(buf)
		// 	fmt.Println(buf.String())
		// }

		fmt.Println(usersArr)

		userJson, err := json.Marshal(usersArr)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			buf := &bytes.Buffer{}
			if err := json.Indent(buf, userJson, "", "\t"); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
			} else {
				c.JSON(200, buf.String())
			}
		}

		// userJson2, err := json.MarshalIndent(usersArr, "", "\t")

		// if err != nil {
		// 	c.JSON(400, gin.H{"error": err.Error()})
		// } else {
		// 	c.JSON(200, string(userJson2))
		// }

	})

	r.POST("/registration", func(c *gin.Context) {
		postFirstname := c.PostForm("firstname")
		postLastname := c.PostForm("lastname")
		postEmail := c.PostForm("email")
		postUsername := c.PostForm("username")

		var token = uuid.New().String()

		c.JSON(200, gin.H{
			"yourToken": token,
			"firstname": postFirstname,
			"lastname":  postLastname,
			"email":     postEmail,
			"username":  postUsername,
		})

		var postUser User = newUser(postUsername, postFirstname, postLastname, postEmail, token)
		// fmt.Println("*-----------------------------------")
		// fmt.Println(postUser)
		// fmt.Println("*-----------------------------------")
		users = append(users, postUser)
		// fmt.Println("**----------------------------------")
		// fmt.Println(users)
		// fmt.Println("**----------------------------------")

	})

	r.POST("/logging", func(c *gin.Context) {
		httpstatus := c.PostForm("httpstatus")
		url := c.PostForm("url")
		application := c.PostForm("application")
		loggingTitle := c.PostForm("title")
		message := c.PostForm("message")
		senderToken := c.PostForm("sender")

		c.JSON(200, gin.H{
			"status":      httpstatus,
			"url":         url,
			"application": application,
			"title":       loggingTitle,
			"message":     message,
			"sender":      senderToken,
		})
	})

	r.Run(":3000") // for a hard coded port

}

func newUser(username string, firstname string, lastname string, email string, token string) User {
	var funcUser = new(User)
	funcUser.Id = uuid.New().String()
	funcUser.Username = username
	funcUser.Firstname = firstname
	funcUser.Lastname = lastname
	funcUser.Email = email
	funcUser.Token = token
	funcUser.LastRequest = nil
	fmt.Println("------------------------------------")
	fmt.Println("New User successfully created:")
	fmt.Println("------------------------------------")
	// fmt.Println(funcUser)
	fmt.Println("\tId\t\t", funcUser.Id, "\n\tUsername\t", funcUser.Username, "\n\tFirstname\t", funcUser.Firstname, "\n\tLastname\t", funcUser.Lastname, "\n\tEmail\t\t", funcUser.Email, "\n\tToken\t\t", funcUser.Token, "\n\tLastRequest\t", funcUser.LastRequest)
	fmt.Println("------------------------------------")

	return *funcUser
}

type User struct {
	Id          string        `json:"id"`
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	Email       string        `json:"email"`
	Username    string        `json:"username"`
	Token       string        `json:"token"`
	LastRequest []lastRequest `json:"lastRequest"`
}

type lastRequest struct {
	id          string    `json:"id"`
	status      string    `json:"status"`
	url         string    `json:"url"`
	application string    `json:"application"`
	title       string    `json:"title"`
	message     string    `json:"message"`
	time        time.Time `json:"time"`
	success     bool      `json:"success"`
}
