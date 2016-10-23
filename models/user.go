package models


import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/oscarvazquez/eleven/db"	
	"time"
	"fmt"
	"log"
	"encoding/json"
	// "reflect" TypeOf
)

type (

	User struct {
		Id 			bson.ObjectId   `json:"id" bson:"_id"`
        Name   		string        	`json:"name" bson:"name"`
        Token  		string          `json:"token" bson:"token"`
        Timestamp	time.Time 	    `json:"created_at" bson: "created_at"`
	}

)

func init(){
	fmt.Println("How many times does this run\n\n\n")
	session := db.GetSession()
	c := session.DB("testEleven").C("users")

    index := mgo.Index{
        Key: []string{"id"},
        Unique: true,
        DropDups: true,
        Background: true,
        Sparse: true,
    }
    err := c.EnsureIndex(index)
    if err != nil {
        log.Fatal(err)
    }
}

func (u User) GetUsers() (string, error) {
	session := db.GetSession()
	defer session.Close()
	c := collection(session)
	var users []User
	
	err := c.Find(bson.M{}).All(&users)
	if err != nil {
        log.Println("Failed get all users: ", err)
        return "", err
	}

    uj, err := json.Marshal(users)
    if err != nil {
    	return "", err
    }

    return string(uj), nil	
}

func (u User) CreateUser() (string, error) {
	session := db.GetSession()
	defer session.Close()
	c := collection(session)

	u.Id = bson.NewObjectId()
	u.Name = "Oscar"
	u.Timestamp = time.Now()

	c.Insert(u)


    uj, err := json.Marshal(u)
    if err != nil {
    	return "", err
    }

    return string(uj), nil
}

func (u *User) ValidateUser(id string) (string, error) {
	session := db.GetSession()
	defer session.Close()
	c := collection(session)
	
	oid := bson.ObjectIdHex(id)
	user := User{}
	err := c.Find(bson.M{"_id": oid}).One(&user)
	
	if err != nil {
		fmt.Println("coudlnt find that user bro")
		return "", err
	}

	if validateTime(user.Timestamp){
		uj, _ := json.Marshal("{success: true}")
		return string(uj), nil
	} else {
		uj, _ := json.Marshal("{success: false}")
		return string(uj), nil
	}
	// return "", nil

}

func validateTime(ac time.Time) (bool){
	fmt.Println(time.Since(ac).Hours())	
	if time.Since(ac).Hours() < 5 {
		return true
	} else {
		return false
	}
}

func collection(session *mgo.Session) (*mgo.Collection){
	return session.DB("testEleven").C("users")
}
