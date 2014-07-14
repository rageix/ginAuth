package main

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/rageix/ginAuth/example/models"
	"github.com/rageix/ginAuth/example/controllers"
	auth "github.com/rageix/ginAuth"
	"fmt"
)

// we load up our database connection suing the Beego ORM (http://beego.me/)
func init() {

	orm.RegisterDriver("postgres", orm.DR_Postgres)

	orm.RegisterDataBase("default", "postgres", "postgres://goserver:goserverpass@localhost/goserver?sslmode=disable")

	orm.RegisterModel(new(models.Users))

}

func main() {

	r := gin.Default()
	r.NotFound404(controllers.Is404)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/login", controllers.LoginPost)
		apiv1.GET("/logout", func(ctx *gin.Context) {
				auth.Logout(ctx) // logs out the user
			})

		authenticate := apiv1.Group("/")  // set up and authentication group
		authenticate.Use(auth.Use)
		authenticate.GET("/checklogin", func(ctx *gin.Context){
				ctx.String(200, "You are logged in!") // if not logged in, you will never reach this
			})

	}

	auth.ConfigPath = "/Users/Brad/go/src/github.com/rageix/ginAuth/example/conf/login.conf"  // the path to our config file
	auth.Unauthorized = controllers.LoginUnauthenticated  // our unauthorized handler

	err := auth.LoadConfig() // load our config file, you can skip this and set the values yourself

	if err != nil {
		fmt.Println(err) // if there was an error loading our config file
	} else {
		r.Run("127.0.0.1:8080") // Run our server if no errors
	}

}
