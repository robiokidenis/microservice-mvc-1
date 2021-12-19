/*
 * Copyright (c) 2021
 * Created at 5/29/21 10:50 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/libraries/httpresponse"
)

func (h *handler) CreateHandler(ctx *gin.Context) {
	var request createRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.Create(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp.GetData())
}
