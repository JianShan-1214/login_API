package main

import (
	"fmt"
	// "io"
	"login-api/check"
	"net/http"
	"net/textproto"
	"time"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username"  binding:"required"`
	Password string `form:"password"  binding:"required"`
}

type Username struct {
	Name string `json:"username" form:"username" binding:"required"`
}

type Userlist struct{
	Name string `json:"username" form:"username" binding:"required"`
	Course []string `json:"course" form:"course" binding:"required"`
}

type Filename struct {
	Name string `json:"filename" form:"filename" binding:"required"`
}

type FileHeader struct {
	Filename string 
	Header textproto.MIMEHeader
	Size int64
	context []byte
	tmpfile string
}

// type File interface {
// 	io.Reader
// 	io.ReaderAt
// 	io.Seeker
// 	io.Closer
// }

func main(){
	r := gin.Default()
	// r.LoadHTMLGlob(("template/html/*"))
	// r.Static("/assets","./template/assets")
	// r.GET("/login",LoginPage)
	// r.POST("/login",LoginAuth)
	r.POST("/login",LoginAuth)
	r_login := r.Group("/",AuthRequired())
	r_login.POST("/logout",Logout)
	r_login.POST("/addUser",Create)
	r_login.POST("/deleteUser",Delete)
	r_login.POST("/userList",List)
	r_login.POST("/uploadVideo",uploadVideo)
	r_login.POST("/rmVideo",Remove)
	r_login.POST("/addCourse",AddCourse)
	r.Run(":8888")

}

// func LoginPage(c *gin.Context){
// 	c.HTML(http.StatusOK,"login.html",nil)
// }

var user User

func AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context){
		cookie,_ :=  c.Request.Cookie("ID")
		fmt.Println(cookie)
		if cookie == nil{
			c.JSON(401,fmt.Sprintf("%s","請登入"))
			c.Abort()
		}
		c.Next()
	}
}

func List(c *gin.Context){
	list := check.UserList()
	c.JSON(200,list)
}

func Logout(c *gin.Context) {
	expires := time.Now().AddDate(0,0,-1)
	cookie := http.Cookie{Name:"ID",Value:"",Expires:expires}
	fmt.Println(cookie)
	http.SetCookie(c.Writer,&cookie)
	c.JSON(200,"logout") 
}

func Create(c *gin.Context){
	c.Bind(&user)
	fmt.Println(user)
	if err :=check.CreateUser(user.Username,user.Password); err != nil{
		c.JSON(401,fmt.Sprintf("%s",err))
	}else{
		c.Redirect(http.StatusMovedPermanently, "./")
		c.JSON(http.StatusOK,"OK")
	}
}

func AddCourse(c *gin.Context){
	var list []Userlist
	c.Bind(&list)
	// fmt.Println(list[0].Course)
	for i := range list{
		if err := check.AddCourse(list[i].Name,list[i].Course);err != nil{
			c.String(401,fmt.Sprintf("%s",err))
		}
	}
	c.String(200,fmt.Sprintf("OK"))
	
}

func LoginAuth(c *gin.Context) {
	c.Bind(&user)
	fmt.Println(user)
	if s,err := check.LoginCheck(user.Username,user.Password);err == nil {
		expires := time.Now().AddDate(0,0,1)
		fmt.Println(expires)
		cookie := http.Cookie{Name:"ID",Value:s,Expires:expires}
		http.SetCookie(c.Writer,&cookie)
		c.JSON(200,"")
	}else{
		c.JSON(401,fmt.Sprintf("%s",err))
	}
}

func Delete(c *gin.Context){
	var name Username
	c.Bind(&name)
	fmt.Println(name)
	if err := check.DeleteUser(name.Name);err == nil {
		c.JSON(200,"OK")
	}else{
		c.JSON(401,fmt.Sprintf("%s",err))
	}
}

func uploadVideo(c *gin.Context){
	file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		if err := c.SaveUploadedFile(file, "./video/"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		fmt.Println(file.Filename)
		check.CreateCourse(file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
}

func Remove(c *gin.Context){
	var file Filename
	c.Bind(&file)
	fmt.Println(file.Name)
	if err:=check.RmCourse(file.Name);err != nil{
		c.String(401,err.Error())
	}else{
		c.String(200,"Success")
	}
}