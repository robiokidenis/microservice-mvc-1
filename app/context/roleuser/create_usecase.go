/*
 * Copyright (c) 2021
 * Created at 5/29/21 10:33 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

import "github.com/robiokidenis/microservice-mvc-1/domain/model"

type CreateRequest interface {
	GetUserID() int64
	GetRoleID() int64
}

type CreateResponse interface {
	GetData() *model.UserRole
}

func (uc *usecase) Create(request CreateRequest) (CreateResponse, error) {
	tx := uc.db.Begin()
	m := new(model.UserRole)
	m.UsersID = request.GetUserID()
	m.RolesID = request.GetRoleID()

	resp, err := uc.roleUserRepo.Create(m, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return createResponse{
		Data: resp,
	}, nil
}
