package main

import (
	// "encoding/json"
	"fmt"
	// "log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string		   `bson:"password"`
	Level    int           `bson:"level"`
}

type List struct {
	Name   string
	UserCourses []string
}

func main() {
	session,err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("CloudClass").C("users")
	// c := session.DB("CloudClass").C("user_class")
	// var list []List
	// err = c.Insert(&User{bson.NewObjectId(),"admin","admin",0})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var user User
	err = c.Find(bson.M{"username": "admin"}).One(&user)
	if err != nil {
		fmt.Println(err)
	}
	// var list []User
	// var all_usernames []string
	// c.Find(bson.M{}).All(&list)
	// fmt.Println(list)
	// for i := range list {
	// 	s := list[i].Username
	// 	all_usernames = append(all_usernames,s)
	// }
	// fmt.Println(all_usernames)
	// userJSON,_ :=json.Marshal(all_usernames)
	// fmt.Println(string(userJSON))

	// var alluser []string
	// for i := range list{
	// 	alluser = append(alluser, list[i].Username)
	// }
	// fmt.Println(alluser)
	fmt.Println(user)
}