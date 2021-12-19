/*
 * Copyright (c) 2021
 * Created at 5/29/21 2:57 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package role

import (
	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/app/context/role"
)

type Handler interface {
	GenerateHandler(ctx *gin.Context)
}

type handler struct {
	uc role.Usecase
}

func NewHandler(uc role.Usecase) Handler {
	return &handler{uc}
}
