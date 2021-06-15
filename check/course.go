package check

import (
	"errors"
	// "fmt"
	"os"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type File struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"filename"`
}

var file File

func CreateCourse(name string){
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("course")
	c.Insert(&File{bson.NewObjectId(),name})
}


func RmCourse(name string)error{
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("course")
	if err :=c.Find(bson.M{"filename":name}).One(&file); err == nil{
		filePath := "./video/"+file.Name
		if err = os.Remove(filePath);err != nil{
			return errors.New("Remove error")
		}else{
			if err = c.RemoveId(file.Id); err != nil{
				return errors.New("Data Error")
			}
		}
		return nil
	}else{
		return errors.New("Not found")
	}
}

func AddCourse(name string, course []string)error{
	session,_ = mgo.Dial("localhost:27017")
	defer session.Close()
	c = session.DB("CloudClass").C("user_course")

	err := c.Update(bson.M{"name":name}, bson.M{"course":course})
	if err != nil{
		return err
	}
	return nil
}