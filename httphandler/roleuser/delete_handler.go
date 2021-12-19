/*
 * Copyright (c) 2021
 * Created at 5/30/21 5:58 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robiokidenis/microservice-mvc-1/libraries/httpresponse"
)

func (h *handler) DeleteHandler(ctx *gin.Context) {
	var request deleteRequest
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	request.ID = int64(ID)

	resp, err := h.uc.Delete(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
