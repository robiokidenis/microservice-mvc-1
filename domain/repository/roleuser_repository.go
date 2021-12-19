/*
 * Copyright (c) 2021
 * Created at 5/29/21 5:23 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
)

type RoleUserRepository interface {
	Create(user *model.UserRole, tx *gorm.DB) (*model.UserRole, error)
	Update(user *model.UserRole, tx *gorm.DB) error
	Delete(user *model.UserRole, tx *gorm.DB) error
	Find(ID int64) (*model.UserRole, error)
	FindOneBy(criteria map[string]interface{}) (*model.UserRole, error)
}
