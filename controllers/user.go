package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/amro-alasri/GoMongoMVCAuthAPI/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	Session *mgo.Session
}

func NewuserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.Session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Parse the JSON request body to extract user information
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing request body: %v", err)
		return
	}

	// create an Id by bson
	newUser.Id = bson.NewObjectId()

	// Insert the new user into the MongoDB "users" collection
	err = uc.Session.DB("mongo-golang").C("users").Insert(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating user: %v", err)
		return
	}

	// Respond with success status
	uj, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
	}
	// Respond with success status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Extract user ID from request parameters
	id := p.ByName("id")

	// Check if the ID is a valid MongoDB ObjectId
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid user ID")
		return
	}

	// Convert the ID to a BSON ObjectId
	oid := bson.ObjectIdHex(id)

	// Remove the user from the MongoDB "users" collection
	err := uc.Session.DB("mongo-golang").C("users").RemoveId(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found: %v", err)
		return
	}

	// Respond with success status
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete user", oid, "\n")
}
