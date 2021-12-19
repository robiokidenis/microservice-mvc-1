/*
 * Copyright (c) 2021
 * Created at 6/2/21 11:07 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package extra

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/libraries/httpresponse"
)

func (h *handler) NotFoundHandler(ctx *gin.Context) {
	url := ctx.Request.RequestURI
	httpresponse.NewErrorResponse(ctx, http.StatusNotFound, errors.New(url+" page not found"))
	return
}
