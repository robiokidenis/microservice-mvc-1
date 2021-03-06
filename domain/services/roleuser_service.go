/*
 * Copyright (c) 2021
 * Created at 5/29/21 5:23 PM
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

type roleUserService struct {
	db *gorm.DB
}

func NewRoleUserService(db *gorm.DB) repository.RoleUserRepository {
	return &roleUserService{db}
}

func (srv *roleUserService) Create(user *model.UserRole, tx *gorm.DB) (m *model.UserRole, err error) {
	db := tx.Create(&user)

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

func (srv *roleUserService) Update(user *model.UserRole, tx *gorm.DB) (err error) {
	err = tx.Save(&user).Error
	return
}

func (srv *roleUserService) Delete(user *model.UserRole, tx *gorm.DB) (err error) {
	err = tx.Delete(&user).Error
	return
}

func (srv *roleUserService) Find(ID int64) (*model.UserRole, error) {
	m := new(model.UserRole)
	row := srv.db.Table("user_role").Select("*").Where("id = ?", ID).Row()
	err := row.Scan(&m.ID, &m.RolesID, &m.UsersID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *roleUserService) FindOneBy(criteria map[string]interface{}) (*model.UserRole, error) {
	m := new(model.UserRole)
	row := srv.db.Table("user_role").Select("*").Where(criteria).Row()
	err := row.Scan(&m.ID, &m.RolesID, &m.UsersID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}
