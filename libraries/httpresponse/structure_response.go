/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:43 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package httpresponse

type ResponseError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
}

type ResponsePaged struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
}

type ResponseObject struct {
	Data interface{} `json:"data"`
}
