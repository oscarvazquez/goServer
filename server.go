package main 

import (
	"fmt"
	"net/http"
	"github.com/oscarvazquez/eleven/config"
	"log"
)
const portnumber = ":8080"

func main(){
	fmt.Printf("Starting server on port %s\n", portnumber)
	router := config.SetRoutes()
	log.Fatal(http.ListenAndServe(portnumber, router))	
}