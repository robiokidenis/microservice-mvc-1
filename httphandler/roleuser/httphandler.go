/*
 * Copyright (c) 2021
 * Created at 5/29/21 10:47 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

import (
	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/app/context/roleuser"
)

type Handler interface {
	CreateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
}

type handler struct {
	uc roleuser.Usecase
}

func NewHandler(uc roleuser.Usecase) Handler {
	return &handler{uc}
}
