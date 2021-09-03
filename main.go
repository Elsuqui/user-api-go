package main

import (
	"UserRestApi/helpers"
	"UserRestApi/migrations"
	"UserRestApi/routes"
	"UserRestApi/services/mysql"
	"UserRestApi/validators"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file to set security or credentials information
	err := godotenv.Load()
	helpers.HandleError(err)
	server := gin.Default()
	db, _ := mysql.Connnect()
	defer mysql.Close(db)
	migrations.Boot(db)
	new(routes.Router).Boot(server)
	validators.NewJSONFormatter()
	server.Run()
}
