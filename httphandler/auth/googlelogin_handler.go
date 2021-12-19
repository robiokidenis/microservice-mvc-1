/*
 * Copyright (c) 2021
 * Created at 5/26/21 1:48 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package auth

import (
	"net/http"

	"github.com/robiokidenis/microservice-mvc-1/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

func (h *handler) GoogleLoginHandle(ctx *gin.Context) {
	request := authRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.GoogleLogin(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
