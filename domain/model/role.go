/*
 * Copyright (c) 2021
 * Created at 5/29/21 2:28 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package model

import "time"

type Role struct {
	ID          int64     `json:"id"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
