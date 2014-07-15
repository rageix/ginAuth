package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/rageix/ginAuth/example/models"
	auth "github.com/rageix/ginAuth"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Email    string
	Password string
}

func LoginPost(ctx *gin.Context) {

	// create a new Beego ORM object
	o := orm.NewOrm()
	o.Using("default")

	// initialize our data structs
	data := LoginData{}
	user := models.Users{}

	// expecting our login data in JSON form
	ctx.EnsureBody(&data)

	// search the users table for the first user with an email and password we supplied
	// of course storing passwords in plaintext is stupid, this is just an example
	err := o.QueryTable("users").Filter("email", data.Email).Filter("password", data.Password).One(&user)

	if err != nil {
		ctx.String(400, "User was not found!")
		return
	}

	// this data will be added to the cookie and available on decode
	extra := map[string]string{"email": data.Email}

	// log in the user
	err1 := auth.Login(ctx, extra)
	if err1 == nil {
		ctx.String(200, "")
	}

}

func LoginAuthenticated(ctx *gin.Context) {

	// load our user or something

}

func LoginUnauthenticated(ctx *gin.Context) {

	ctx.String(401, "You are not logged in!")

	// ctx.Abort() means no other handlers will be ran after this one, meaning the route controller won't execute
	ctx.Abort(-1)

}
