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

func result(ctx *gin.Context, status int, message interface{}) {

	ctx.JSON(status, message)
	ctx.Abort(-1) // stops all execution of future handlers and returns our response

}

func LoginPost(ctx *gin.Context) {

	o := orm.NewOrm()
	o.Using("default")

	data := LoginData{}
	user := models.Users{}

	ctx.EnsureBody(&data)

	err := o.QueryTable("users").Filter("email", data.Email).Filter("password", data.Password).One(&user)

	if err != nil {
		result(ctx, 400, "User was not found!")
		return
	}

	extra := map[string]string{"email": data.Email}  // this data will be added to the cookie and available on decode

	err1 := auth.Login(ctx, extra)
	if err1 == nil {
		result(ctx, 200, "")
	}

}

func LoginAuthenticated(ctx *gin.Context) {

	// load our user or something

}

func LoginUnauthenticated(ctx *gin.Context) {

	result(ctx, 401, "You are not logged in!")

}
