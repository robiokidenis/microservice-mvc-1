/*
 * Copyright (c) 2021
 * Created at 6/2/21 11:06 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package extra

import "github.com/gin-gonic/gin"

type Handler interface {
	NotFoundHandler(ctx *gin.Context)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}
