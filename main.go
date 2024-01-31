package main

import (
	"colls.labs.bake/database"
	"colls.labs.bake/items"
	"colls.labs.bake/orders"

	"colls.labs.bake/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// models migration
func Migration(database *gorm.DB) {

	users.AutoMigrate()
	items.AutoMigrate()
	orders.AutoMigrate()
}

func main() {
	// Initialize the database connection
	database := database.Init()
	Migration(database)

	router := gin.Default()

	// CORS configuration :updated on august 28 17:16 pm
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	sun := router.Group("/Api")
	users.Create(sun.Group("/users"))
	items.Create(sun.Group("/items"))
	orders.Create(sun.Group("/orders"))
	if err := router.Run(":8082"); err != nil {
		panic(err)
	}
}
