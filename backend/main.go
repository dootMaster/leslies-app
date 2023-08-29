package main

import (
	"log"
	"net/http"
	"os"

	ctrl "leslies-app/backend/controllers"
	"leslies-app/backend/db"
	mw "leslies-app/backend/middleware"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		panic("Error loading .env file")
	}
	log.SetOutput(os.Stdout)

	db.ConnectToDB()
	defer db.CloseDBConnection()

	http.HandleFunc("/", mw.Auth(mw.HttpMethod(mw.GET, ctrl.HealthCheck)))
	http.HandleFunc("/user/create", mw.HttpMethod(mw.POST, ctrl.CreateUser))
	http.HandleFunc("/user/login", mw.HttpMethod(mw.POST, ctrl.Login))
	http.HandleFunc("/user/logout", mw.HttpMethod(mw.POST, ctrl.Logout))

	port := os.Getenv("SERVER_PORT")
	log.Printf("Server started on port %s.\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Print("Server failed to start.")
		panic(err)
	}
}
