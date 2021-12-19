/*
 * Copyright (c) 2021
 * Created at 5/28/21 12:09 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package auth

import (
	"errors"

	"github.com/robiokidenis/microservice-mvc-1/app/config"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
	"github.com/robiokidenis/microservice-mvc-1/libraries/helper"
	"github.com/robiokidenis/microservice-mvc-1/libraries/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) PhoneLogin(request PhoneLoginRequest) (AuthResponse, error) {

	phone := helper.NewPhoneNumber(request.GetPhone())
	user, err := uc.getUserByPhone(phone.GetPhoneNumber())

	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.GetPin())); err != nil {
		return nil, errors.New("phone and password not match")
	}

	var role string
	roles, err := uc.findRoleUser(user)
	if err != nil {
		role = config.CONSUMER
	} else {
		role = roles.Value
	}

	/** generate parameter for Custom JWT */
	param := middleware.TokenStructure{}
	param.UserID = user.ID
	param.Email = user.Email
	param.Role = role

	auth := middleware.NewCustomAuth([]byte(config.SIGNATURE))
	token, err := auth.GenerateToken(param)
	if err != nil {
		return nil, err
	}

	if err = uc.updateLastLogin(user); err != nil {
		return nil, err
	}

	return token, nil
}

func (uc *usecase) getUserByPhone(phone string) (*model.Users, error) {
	criteria := make(map[string]interface{})
	criteria["phone"] = phone
	criteria["is_verified"] = true

	user, err := uc.userRepo.FindOneBy(criteria)
	if err != nil {
		return nil, err
	}

	return user, nil
}
