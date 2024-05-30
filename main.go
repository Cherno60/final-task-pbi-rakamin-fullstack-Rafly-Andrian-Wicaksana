package main

import (
	"github.com/gin-gonic/gin"
	controller "goAPI/Controller"
	database "goAPI/Database"
	router "goAPI/Router"
	"net/http"
	"os"
)

func init() {
	database.LoadEnvVariables()
	database.DBConnect()
}

func main() {

	//router.GetRouter("register", "register")
	//router.GetRouter("home", "none")
	//router.Run("localhost:9090")
	r := gin.Default()
	r.GET("/home", router.Hello)
	r.POST("/login", controller.Login)
	r.GET("/photos", controller.PhotosIndex)
	r.GET("/register", router.Register)
	r.Run("localhost:" + os.Getenv("PORT"))
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
