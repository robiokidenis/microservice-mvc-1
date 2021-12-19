/*
 * Copyright (c) 2021
 * Created at 5/30/21 5:52 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package roleuser

type DeleteResponse interface {
}

type DeleteRequest interface {
	GetID() int64
}

func (uc *usecase) Delete(request DeleteRequest) (DeleteResponse, error) {
	m, err := uc.roleUserRepo.Find(request.GetID())
	if err != nil {
		return nil, err
	}

	tx := uc.db.Begin()
	if err = uc.roleUserRepo.Delete(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return deleteResponse{
		Message: "Success",
	}, nil
}
