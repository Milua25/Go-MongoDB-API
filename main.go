package main

import (
	"github.com/Golang-Personal-Projects/Go-Projects/04-Go-MongoDB-API/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	// create a new router
	router := httprouter.New()
	userController := controllers.NewUserController(getSession())

	router.GET("/user/:id", userController.GetUser)
	router.DELETE("/user/:id", userController.DeleteUser)
	router.POST("/user/", userController.CreateUser)

	// Listen and serve
	if err := http.ListenAndServe(":8080", router); err == nil {
		log.Fatal("Server failed to start")
	}
}

// connect to the MongoDB
func getSession() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		log.Fatal(err)
	}
	return session
}
