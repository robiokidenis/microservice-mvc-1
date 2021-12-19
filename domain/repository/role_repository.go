/*
 * Copyright (c) 2021
 * Created at 5/29/21 2:29 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
)

type RoleRepository interface {
	Create(role *model.Role, tx *gorm.DB) (*model.Role, error)
	FindOneBy(criteria map[string]interface{}) (*model.Role, error)
	Find(ID int64) (*model.Role, error)
}
