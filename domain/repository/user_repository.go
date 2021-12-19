/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:27 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
)

type UserRepository interface {
	Create(users *model.Users, tx *gorm.DB) (*model.Users, error)
	Update(users *model.Users) (err error)
	Find(ID int64) (*model.Users, error)
	FindOneBy(map[string]interface{}) (*model.Users, error)
}
