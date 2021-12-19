/*
 * Copyright (c) 2021
 * Created at 5/29/21 10:48 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

type createRequest struct {
	RoleID int64 `json:"role_id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

func (req createRequest) GetRoleID() int64 {
	return req.RoleID
}

func (req createRequest) GetUserID() int64 {
	return req.UserID
}
