/*
 * Copyright (c) 2021
 * Created at 5/29/21 3:36 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package role

import "github.com/robiokidenis/microservice-mvc-1/domain/model"

type generateResponse struct {
	Data []*model.Role
}

func (resp generateResponse) GetData() []*model.Role {
	return resp.Data
}
