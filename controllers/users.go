package controllers

import (
	"net/http"
	"github.com/oscarvazquez/eleven/models"
	"github.com/gorilla/mux"	
	"fmt"
	"log"	
)

type (  
    UserController struct{}
)

func NewUserController() *UserController {  
    return &UserController{}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	u := models.User{}
	uj, err := u.CreateUser();
	if(err != nil){
        log.Fatal(err)		
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request){
	u := models.User{}
	
	uj, err := u.GetUsers();
	if(err != nil){
        log.Fatal(err)		
	}
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", uj)	
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request){
	u := models.User{}
	uj, err := u.ValidateUser(mux.Vars(r)["id"]);
	if(err != nil){
		log.Fatal(err)
	}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", uj)
}