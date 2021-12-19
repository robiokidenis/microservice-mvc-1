/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:18 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMysqlConfig(cfg Config) (err error, db *gorm.DB) {
	dbUser := cfg.GetString(`mysql.user`)
	dbPass := cfg.GetString(`mysql.pass`)
	dbName := cfg.GetString(`mysql.database`)
	dbHost := cfg.GetString(`mysql.address`)
	dbPort := cfg.GetString(`mysql.port`)

	db, err = gorm.Open("mysql", ""+dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed connect to database")
	}

	return err, db
}
