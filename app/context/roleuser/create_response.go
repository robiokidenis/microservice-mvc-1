/*
 * Copyright (c) 2021
 * Created at 5/29/21 10:39 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

import "github.com/robiokidenis/microservice-mvc-1/domain/model"

type createResponse struct {
	Data *model.UserRole
}

func (resp createResponse) GetData() *model.UserRole {
	return resp.Data
}
