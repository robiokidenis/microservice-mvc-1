/*
 * Copyright (c) 2021
 * Created at 5/29/21 3:05 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/libraries/httpresponse"
)

func (h *handler) GenerateHandler(ctx *gin.Context) {
	resp, err := h.uc.Generate()
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp.GetData())
}
