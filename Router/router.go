package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func GetRouter(route, payload string) {
//	router := gin.Default()
//
//	//result := fmt.Println(hello())
//	if payload == "none" {
//		router.GET("/"+route, Hello)
//
//	} else if payload == "register" {
//		router.GET("/"+route, Register)
//	}
//	router.Run("localhost:9090")
//
//}

type user struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Welcome to home page")
}

func Login(c *gin.Context) {
	var payload loginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin := "Unauthorized"
	if payload.Username == "admin" && payload.Password == "123456" {
		admin = "Welcome Admin"
	}
	c.JSON(http.StatusOK, gin.H{"message": admin})
}

func Register(c *gin.Context) {
	c.String(http.StatusOK, "regis")
}

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{
		ID:        "1",
		Item:      "Clean Room",
		Completed: true,
	},
	{
		ID:        "2",
		Item:      "Read Book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Code",
		Completed: false,
	},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
