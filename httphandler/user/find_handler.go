/*
 * Copyright (c) 2021
 * Created at 5/27/21 1:34 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package user

import (
	"net/http"

	"github.com/robiokidenis/microservice-mvc-1/libraries/helper"
	"github.com/robiokidenis/microservice-mvc-1/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

func (h *handler) FindHandler(ctx *gin.Context) {
	userID, err := helper.GetAuthenticatedUser(ctx.Request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.Find(userID)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
