/*
 * Copyright (c) 2021
 * Created at 5/27/21 1:26 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package user

import "github.com/robiokidenis/microservice-mvc-1/domain/repository"

type Usecase interface {
	Find(ID int64) (FindResponse, error)
}

type usecase struct {
	userRepo repository.UserRepository
}

func NewUsecase(userRepo repository.UserRepository) Usecase {
	return &usecase{userRepo}
}
