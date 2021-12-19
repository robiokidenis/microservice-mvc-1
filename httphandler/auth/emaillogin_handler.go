/*
 * Copyright (c) 2021
 * Created at 5/28/21 12:59 PM
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

func (h *handler) EmailLoginHandler(ctx *gin.Context) {
	request := emailLoginRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.EmailLogin(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
