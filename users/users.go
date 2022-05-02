package users

import (
	"bookspreadLogging/requests"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var users []User

func BasicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	fmt.Println(user, password, hasAuth)
	fmt.Println(users)
	if hasAuth {
		i := sort.Search(len(users), func(i int) bool { return user <= users[i].Username })
		if i < len(users) && users[i].Username == user {
			if checkPasswordHash(password, users[i].Password) {
				// c.JSON(200, gin.H{"message": "user found"})
				fmt.Println("successfully")
			} else {
				// c.JSON(400, gin.H{"error": "password is not correct"})
				c.Abort()
				c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
				c.JSON(401, gin.H{"error": "unauthorized"})
			}
		} else {
			c.Abort()
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			c.JSON(400, gin.H{"error": "user not found"})
		}
	}
}

func Registration(c *gin.Context) {
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	username := c.PostForm("username")
	email := c.PostForm("email")
	password, _ := hashPassword(c.PostForm("password"))
	role := c.PostForm("role")
	currentUser := newUser(firstname, lastname, username, email, password, role)
	users = append(users, currentUser)
	c.JSON(200, currentUser)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUsers(c *gin.Context) {
	if users != nil {
		c.JSON(200, users)
	} else {
		c.JSON(400, gin.H{"error": "No users found"})
	}
}

func newUser(firstname string, lastname string, username string, email string, password string, role string) User {
	var currentUser = new(User)
	currentUser.Id = uuid.New().String()
	currentUser.Firstname = firstname
	currentUser.Lastname = lastname
	currentUser.Username = username
	currentUser.Email = email
	currentUser.Password = password
	currentUser.Role = role

	return *currentUser
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	i := sort.Search(len(users), func(i int) bool { return id <= users[i].Id })
	if i < len(users) && users[i].Id == id {
		users = append(users[:i], users[i+1:]...)
		c.JSON(200, gin.H{"message": "delete user with the id: " + id})
	} else {
		c.JSON(400, gin.H{"error": "No user found with the id: " + id})
	}
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	i := sort.Search(len(users), func(i int) bool { return id <= users[i].Id })
	if i < len(users) && users[i].Id == id {
		c.JSON(200, users[i])
	} else {
		c.JSON(400, gin.H{"error": "No user found with the id: " + id})
	}
}

func EditUser(c *gin.Context) {
	id := c.Param("id")

	i := sort.Search(len(users), func(i int) bool { return id <= users[i].Id })
	if i < len(users) && users[i].Id == id {
		users[i].Firstname = c.PostForm("firstname")
		users[i].Lastname = c.PostForm("lastname")
		users[i].Username = c.PostForm("username")
		users[i].Email = c.PostForm("email")
		users[i].Password, _ = hashPassword(c.PostForm("password"))
		users[i].Role = c.PostForm("role")
		c.JSON(200, users[i])
	} else {
		c.JSON(400, gin.H{"error": "No user found with the id: " + id})
	}
}

// For Requst

func CreateRequst(c *gin.Context) {
	currentUsername, _, _ := c.Request.BasicAuth()
	request := requests.NewRequest(c)

	i := sort.Search(len(users), func(i int) bool { return currentUsername <= users[i].Username })
	if i < len(users) && users[i].Username == currentUsername {
		users[i].Request = append(users[i].Request, request)
		c.JSON(200, users[i])
	}

}

type User struct {
	Id        string
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
	Role      string
	Request   []requests.Request
}
