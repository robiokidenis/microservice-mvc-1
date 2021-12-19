/*
 * Copyright (c) 2021
 * Created at 5/20/21 11:02 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package router

import (
	"github.com/robiokidenis/microservice-mvc-1/app/config"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/auth"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/role"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/roleuser"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/user"
	"github.com/robiokidenis/microservice-mvc-1/libraries/middleware"
	"github.com/gin-gonic/gin"
)

func InvokeRoute(
	engine *gin.Engine,
	auth auth.Handler,
	user user.Handler,
	role role.Handler,
	roleuser roleuser.Handler,
) {

	engine.NoRoute()

	conf := config.NewConfig()

	markone := engine.Group("api/" + conf.GetString("app.version.tag") + conf.GetString("app.version.value"))
	markone.Use(gin.Logger())
	markone.Use(gin.Recovery())
	markone.Use(gin.ErrorLogger())

	/** auth route group */
	{
		authroute := markone.Group("auth")
		authroute.POST("/google", auth.GoogleLoginHandle)
		authroute.POST("/facebook", auth.FacebookLoginHandle)
		authroute.POST("/phone", auth.PhoneLoginHandler)
		authroute.POST("/email", auth.EmailLoginHandler)
	}

	/** profile route group */
	{
		me := markone.Group("me")
		me.Use(middleware.MyAuth())
		me.GET("", user.FindHandler)
	}

	/** role route group */
	{
		roleroute := markone.Group("roles")
		roleroute.GET("/generate", role.GenerateHandler)
	}

	/** roleuser route group */
	{
		roleuserroute := markone.Group("roleuser")
		roleuserroute.Use(middleware.MyAuth(config.ADMIN))
		roleuserroute.POST("", roleuser.CreateHandler)
		roleuserroute.DELETE(":id", roleuser.DeleteHandler)
	}

}
