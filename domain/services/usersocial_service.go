/*
 * Copyright (c) 2021
 * Created at 5/26/21 11:15 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package services

import (
	"encoding/json"

	"github.com/robiokidenis/microservice-mvc-1/domain/model"
	"github.com/robiokidenis/microservice-mvc-1/domain/repository"
	"github.com/jinzhu/gorm"
)

type userSocialServices struct {
	db *gorm.DB
}

func NewUserSocialServices(db *gorm.DB) repository.UserSocialRepository {
	return &userSocialServices{db}
}

func (srv *userSocialServices) FindOneBy(criteria map[string]interface{}) (*model.UserSocial, error) {
	m := new(model.UserSocial)
	row := srv.db.Table("user_socials").Select("*").Where(criteria).Row()
	if err := row.Scan(&m.ID, &m.UsersID, &m.SocialName, &m.SocialID, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *userSocialServices) Create(social *model.UserSocial, tx *gorm.DB) (m *model.UserSocial, err error) {
	db := tx.Create(&social)
	if err = db.Error; err != nil {
		return
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return
	}
	if err = json.Unmarshal(byteData, &m); err != nil {
		return
	}

	return
}
