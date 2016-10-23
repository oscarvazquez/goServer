package config

import (
	"github.com/gorilla/mux"
	"github.com/oscarvazquez/eleven/controllers"

)

func SetRoutes() *mux.Router{
	uc := controllers.NewUserController()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", uc.CreateUser)
	myRouter.HandleFunc("/users/index", uc.GetUsers)
	myRouter.HandleFunc("/users/{id}", uc.Validate)	
	return myRouter
}
