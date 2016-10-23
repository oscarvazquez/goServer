package db

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func init() {
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