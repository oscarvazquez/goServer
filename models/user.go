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
		Uid			string			`json:"uid" bson:"_uid"` // this will be steam id
        Name   		string        	`json:"name" bson:"name"`
        Token  		string          `json:"token" bson:"token"` 
        Timestamp	time.Time 	    `json:"created_at" bson: "created_at"`
        Validated 	bool 			`json:"validated" bson: "validated"`
	}

	Validation struct {
		Username 	string   		`json:"username" bson:"username"`		
		Message		string 	 		`json:"message"`
		Status		bool 	 		`json:"status"`
	}

)

const (
	maxDemoTime 	=   (604800)
	daysCalc 		= 	(60 * 60 * 24)
	hoursCalc 		=   (60 * 60)
	minutesCalc 	= 	(60)
)

func init(){
	setIndices("id")
	setIndices("uid")
}

func setIndices(s string){
	session := db.GetSession()
	defer session.Close()

	c := collection(session)	
    index := mgo.Index{
        Key: []string{s},
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

	return packageUsers(users)	
}

func (u User) CreateUser() (string, error) {
	session := db.GetSession()
	defer session.Close()

	c := collection(session)

	u.Id = bson.NewObjectId()
	u.Name = "Oscar"
	u.Timestamp = time.Now()

	c.Insert(u)

	return packageUser(u)
}

func (u *User) ValidateUser(id string) (string, error) {
	session := db.GetSession()
	defer session.Close()

	c := collection(session)

	oid := bson.ObjectIdHex(id)
	user := User{}
	err := c.Find(bson.M{"_id": oid}).One(&user)
	
	if err != nil {
		fmt.Println("could not find that user bro")
		return "", err
	}
	var v Validation
	v.validateTime(user)
	return packageValidation(v)
}

func packageUser(u User) (string, error){	
    uj, err := json.Marshal(u)
    if err != nil {
    	return "", err
    }

    return string(uj), nil	
}

func packageUsers(users []User) (string, error){
    uj, err := json.Marshal(users)
    if err != nil {
    	return "", err
    }

    return string(uj), nil		
}

func (v *Validation) validateTime(u User){
	var timeLeft = maxDemoTime - int(time.Since(u.Timestamp).Seconds())
	if timeLeft > 0 {
		v.Status = true
		days := int(timeLeft / daysCalc)
		timeLeft -= int(days * daysCalc)
		hours := int(timeLeft / hoursCalc)
		timeLeft -= int(hours * hoursCalc)
		minutes := int(timeLeft / minutesCalc)
		v.Message = fmt.Sprintf("You still have %d days, %d hours and %d minutes left to play", days, hours, minutes)
	} else {
		v.Status = false
		v.Message = "You ran out of time"
	}
	v.Username = u.Name
}

func packageValidation(v Validation) (string, error){	
	vm, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(vm), nil
}

func collection(session *mgo.Session) (*mgo.Collection){
	return session.DB("testEleven").C("users")
}
