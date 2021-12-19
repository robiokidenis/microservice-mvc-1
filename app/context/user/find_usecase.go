/*
 * Copyright (c) 2021
 * Created at 5/27/21 1:28 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package user

type FindResponse interface {
}

func (uc *usecase) Find(ID int64) (FindResponse, error) {
	users, err := uc.userRepo.Find(ID)
	if err != nil {
		return nil, err
	}

	return users, nil
}
