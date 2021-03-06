package router

import (
	"fmt"
	"log"

	"spa-go/cmd/app/handlers/base"
	v1 "spa-go/cmd/app/handlers/v1"
	middleware "spa-go/internal/mid"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	r := setupRouter()
	port := fmt.Sprintf(":%s", viper.GetString("app.port"))
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}

func setMode() {
	switch viper.GetString("app.runMode") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		log.Println("Load App Mode Error!")
	}
}

func setupRouter() *gin.Engine {
	setMode()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.JWT())
	r.Use(middleware.Cors())
	r.GET("/health", base.Health)

	api := r.Group("/api/v1")

	user := api.Group("/user")
	user.Use(middleware.JWT())
	{
		user.POST("", v1.UserCreate)
		user.GET("/:id", v1.UserGet)
		user.PUT("/:id", v1.UserUpdate)
		user.DELETE("/:id", v1.UserDelete)
		// user.GET("/", v1.UserList)
	}

	// policy := api.Group("/policy")
	// policy.Use(middleware.Casbin(auth.Casbin))

	// {
	// 	policy.POST("", v1.PolicyCreate)
	// 	policy.GET("/:id", v1.PolicyGet)
	// }

	login := api.Group("/admin")
	{
		login.POST("", v1.LogIn)
		login.PUT("", v1.UpdatePasswd)
	}

	return r
}
