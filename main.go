package main

import (
	"4mount/backend/controllers"
	"4mount/backend/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("start")
	initializers.LoadEnvInitializers()
	initializers.ConnectToDB()
	fmt.Println("end")
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.Run()
}
