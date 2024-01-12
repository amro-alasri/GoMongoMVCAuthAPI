package main

import (
	"github.com/amro-alasri/GoMongoMVCAuthAPI/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewuserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8040", r)
}

func getSession() *mgo.Session {
	// Connect to MongoDB
	session, err := mgo.Dial("mongodb://localhost:27017") // Replace with your MongoDB connection string
	if err != nil {
		panic(err)
	}
	return session
}
