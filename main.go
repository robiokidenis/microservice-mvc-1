/*
 * Copyright (c) 2021
 * Created at 5/28/21 7:32 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package main

import (
	"fmt"

	"github.com/robiokidenis/microservice-mvc-1/app/config"
	"github.com/robiokidenis/microservice-mvc-1/libraries/middleware"
	"github.com/robiokidenis/microservice-mvc-1/router"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

var _ = dig.Name

func main() {
	var err error

	myrole := make(map[string][]string)
	myrole[config.ADMIN] = []string{config.ADMIN}
	myrole[config.MERCHANT] = []string{config.ADMIN, config.MERCHANT}
	myrole[config.CONSUMER] = []string{config.ADMIN, config.MERCHANT, config.CONSUMER}

	middleware.InitRole(myrole)
	middleware.InitJWTMiddlewareCustom([]byte(config.SIGNATURE), jwt.SigningMethodHS512)

	gin.SetMode(gin.DebugMode)

	c := NewBuildContainer()

	if err = c.Invoke(router.InvokeRoute); err != nil {
		panic(err)
	}

	if err = c.Provide(NewRoute); err != nil {
		panic(err)
	}

	if err = c.Invoke(func(s *ServerRoute) {
		s.Run()
	}); err != nil {
		fmt.Println(err)
	}
}
