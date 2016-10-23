package main 

import (
	"fmt"
	"net/http"
	"github.com/oscarvazquez/eleven/config"
	// "github.com/oscarvazquez/eleven/db"	
	"log"
)
const portnumber = ":8080"

func main(){
	fmt.Printf("Starting server on port %s\n", portnumber)
	// db.CreateSession()
	router := config.SetRoutes()
	log.Fatal(http.ListenAndServe(portnumber, router))	
}