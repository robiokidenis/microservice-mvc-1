/*
 * Copyright (c) 2021
 * Created at 5/28/21 12:10 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package auth

type phoneLoginRequest struct {
	Phone string `json:"phone" binding:"required"`
	Pin   string `json:"pin" binding:"required"`
}

func (req phoneLoginRequest) GetPhone() string {
	return req.Phone
}

func (req phoneLoginRequest) GetPin() string {
	return req.Pin
}
