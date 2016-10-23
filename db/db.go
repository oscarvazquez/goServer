package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

var session *mgo.Session

func init() {
	fmt.Println("Creating my original session\n\n\n")

    s, err := mgo.Dial("mongodb://localhost/")
    if err != nil {
        panic(err)
    }
    s.SetMode(mgo.Monotonic, true)

 	session = s
}


func GetSession() *mgo.Session {
	return session.Copy()
}