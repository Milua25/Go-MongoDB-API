package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Golang-Personal-Projects/Go-Projects/04-Go-MongoDB-API/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	// return status 200
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)

	// create a new object ID
	u.Id = bson.NewObjectId()

	// Insert the new user
	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)
	// check for errors
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	// return status 201
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// return status 201
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User", oid, "\n")
}
