/*
 * Copyright (c) 2021
 * Created at 5/30/21 5:57 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

type deleteRequest struct {
	ID int64 `json:"id"`
}

func (req deleteRequest) GetID() int64 {
	return req.ID
}
