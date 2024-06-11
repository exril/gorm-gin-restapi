package main

import (
	"4mount/backend/initializers"
	"4mount/backend/models"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}