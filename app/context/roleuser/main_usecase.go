/*
 * Copyright (c) 2021
 * Created at 5/29/21 5:19 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

import (
	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/repository"
)

type Usecase interface {
	Create(request CreateRequest) (CreateResponse, error)
	Delete(request DeleteRequest) (DeleteResponse, error)
}

type usecase struct {
	roleUserRepo repository.RoleUserRepository
	db           *gorm.DB
}

func NewUsecase(roleUserRepo repository.RoleUserRepository, db *gorm.DB) Usecase {
	return &usecase{roleUserRepo, db}
}
