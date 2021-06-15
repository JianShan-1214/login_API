package check

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string		   `bson:"password"`
	Level    int           `bson:"level"`
}

type Userlist struct{
	Name string `json:"username" form:"username" binding:"required"`
	Course []string `json:"course" form:"course" binding:"required"`
}

var c *mgo.Collection
var session *mgo.Session
var user User


func CreateUser(name string,ps string)error{
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("users")
	c_c := session.DB("CloudClass").C("user_class")
	var user User
	if err :=c.Find(bson.M{"username":name}).One(&user); err != nil{
		c.Insert(&User{bson.NewObjectId(),name,ps,1})
		c_c.Insert(&Userlist{name,[]string{}})
	}else{
		return errors.New("已有此帳號")
	}
	return nil
}

func UserList()(all_user []Userlist){
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("users")
	var list []User
	s := []string{"111","222"}
	c.Find(bson.M{}).All(&list)
	for i := range list {
		temp := Userlist{list[i].Username,s}
		all_user = append(all_user,temp)
	}

	return 
}

func LoginCheck(name string,ps string)(hex string,err error){
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("users")
	if err :=c.Find(bson.M{"username":name}).One(&user); err == nil{
		if ps==user.Password{
			return user.Id.Hex(),nil
		}
	}
	return "",errors.New("Error")
}

func DeleteUser(name string)error{
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("users")
	err := c.Find(bson.M{"username":name}).One(&user)
	fmt.Println(user)
	if err == nil && user.Username != "admin"{
		if err = c.RemoveId(user.Id); err != nil{
			return errors.New("Data Error")
		}
		return nil
	}
	return errors.New("Not found")
}


