/*
 * Copyright (c) 2021
 * Created at 5/29/21 2:41 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package role

import (
	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
	"github.com/robiokidenis/microservice-mvc-1/domain/repository"
)

type Usecase interface {
	Generate() (GenerateResponse, error)
}

type usecase struct {
	roleRepo repository.RoleRepository
	db       *gorm.DB
}

func NewUsecase(roleRepo repository.RoleRepository,
	db *gorm.DB) Usecase {
	return &usecase{roleRepo, db}
}

func (uc *usecase) findByValue(role string) (roles *model.Role, err error) {
	criteria := make(map[string]interface{})
	criteria["value"] = role

	roles, err = uc.roleRepo.FindOneBy(criteria)
	return
}

func (uc *usecase) create(roles *model.Role, tx *gorm.DB) (*model.Role, error) {
	role, err := uc.roleRepo.Create(roles, tx)
	if err != nil {
		return nil, err
	}
	return role, nil
}
