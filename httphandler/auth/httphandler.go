/*
 * Copyright (c) 2021
 * Created at 5/26/21 1:44 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/app/context/auth"
)

type Handler interface {
	GoogleLoginHandle(ctx *gin.Context)
	FacebookLoginHandle(ctx *gin.Context)
	PhoneLoginHandler(ctx *gin.Context)
	EmailLoginHandler(ctx *gin.Context)
}

type handler struct {
	uc auth.Usecase
}

func NewHandler(uc auth.Usecase) Handler {
	return &handler{uc}
}
