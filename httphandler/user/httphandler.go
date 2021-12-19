/*
 * Copyright (c) 2021
 * Created at 5/27/21 1:33 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/app/context/user"
)

type Handler interface {
	FindHandler(ctx *gin.Context)
}

type handler struct {
	uc user.Usecase
}

func NewHandler(uc user.Usecase) Handler {
	return &handler{uc}
}
