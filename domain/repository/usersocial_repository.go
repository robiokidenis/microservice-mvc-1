/*
 * Copyright (c) 2021
 * Created at 5/26/21 11:14 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
)

type UserSocialRepository interface {
	FindOneBy(criteria map[string]interface{}) (*model.UserSocial, error)
	Create(social *model.UserSocial, tx *gorm.DB) (*model.UserSocial, error)
}
