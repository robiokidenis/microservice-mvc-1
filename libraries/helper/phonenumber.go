/*
 * Copyright (c) 2021
 * Created at 5/28/21 12:16 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package helper

import "strings"

type PhoneNumber interface {
	GetPhoneNumber() string
	Sanitize() string
}

type phoneNumber struct {
	phone string
}

func NewPhoneNumber(phone string) PhoneNumber {
	return &phoneNumber{phone}
}

func (this phoneNumber) GetPhoneNumber() string {
	phone := this.Sanitize()
	return "62" + phone[1:]
}

func (this phoneNumber) Sanitize() string {
	phone := strings.Replace(this.phone, " ", "", 0)
	phone = strings.Replace(phone, "-", "", 0)
	if phone[0:2] == "62" {
		return "0" + phone[2:]
	} else if phone[0:3] == "+62" {
		return "0" + phone[3:]
	}
	return phone
}
