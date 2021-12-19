/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:26 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/robiokidenis/microservice-mvc-1/domain/model"
	"github.com/robiokidenis/microservice-mvc-1/domain/repository"
)

type userServices struct {
	db *gorm.DB
}

func NewUserServices(db *gorm.DB) repository.UserRepository {
	return &userServices{db}
}

func (srv *userServices) Create(users *model.Users, tx *gorm.DB) (m *model.Users, err error) {
	db := tx.Create(&users)

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

func (srv *userServices) Update(users *model.Users) (err error) {
	err = srv.db.Save(&users).Error
	return
}

func (srv *userServices) Find(ID int64) (*model.Users, error) {
	m := new(model.Users)
	res := srv.db.Where("id=?", ID).Find(&m).Related(&m.Socials).Related(&m.RoleUsers).Scan(&m)
	if err := res.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (srv *userServices) FindOneBy(criteria map[string]interface{}) (*model.Users, error) {
	m := new(model.Users)
	row := srv.db.Table("users").Select("*").Where(criteria).Row()
	if err := row.Scan(&m.ID, &m.Name, &m.Username, &m.Email, &m.Phone, &m.Password, &m.LastLogin, &m.SecondaryEmail,
		&m.IsVerified, &m.Avatar, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return nil, err
	}
	return m, nil
}
