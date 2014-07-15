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

// we load up our database connection using the Beego ORM (http://beego.me/docs/mvc/model/overview.md)
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

				// logs out the user
				auth.Logout(ctx)

			})

		// set up and authentication group that uses our ginAuth middleware
		authenticate := apiv1.Group("/")
		// tell it to use the ginAuth middleware on these routes
		authenticate.Use(auth.Use)
		authenticate.GET("/checklogin", func(ctx *gin.Context){

				// if not logged in, you will never reach this
				ctx.String(200, "You are logged in!")

			})

	}

	// the path to our config file
	auth.ConfigPath = "/Users/Brad/go/src/github.com/rageix/ginAuth/example/conf/login.conf"
	// our unauthorized handler
	auth.Unauthorized = controllers.LoginUnauthenticated
	// our authorized handler
	auth.Authorized = controllers.LoginAuthenticated

	// load our config file, you can skip this and set the values yourself
	err := auth.LoadConfig()

	if err != nil {

		// if there was an error loading our config file
		fmt.Println(err)

	} else {

		// Run our server if no errors
		r.Run("127.0.0.1:8080")

	}

}
