package main

import (
	"api-labstats/controllers"
	"api-labstats/repository"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

func main() {
	//Load .env
	loadEnv()

	//Connect to DB
	repository.Init()

	//Set timezone
	var errLoc error
	time.Local, errLoc = time.LoadLocation("America/Caracas")
	if errLoc != nil {
		log.Printf("error loading location %v\n", errLoc)
	}

	//Start Iris Server
	app := iris.New()
	api := app.Party("/api/v1")
	{
		auth := api.Party("/auth")
		{
			auth.Post("/login", controllers.AuthHandler)
		}
	}
	app.Listen(":8080")

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
