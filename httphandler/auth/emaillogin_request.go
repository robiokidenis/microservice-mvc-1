/*
 * Copyright (c) 2021
 * Created at 5/28/21 12:58 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package auth

type emailLoginRequest struct {
	Email string `json:"email" binding:"required"`
	Pin   string `json:"pin" binding:"required"`
}

func (req emailLoginRequest) GetEmail() string {
	return req.Email
}

func (req emailLoginRequest) GetPin() string {
	return req.Pin
}
