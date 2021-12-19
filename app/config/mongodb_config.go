/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:21 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

func NewMongoDBConfig(cfg Config) (error, *mgo.Database) {
	address := cfg.GetString(`mongodb.address`)
	database := cfg.GetString(`mongodb.database`)
	log.Println("mongo config", address, database)
	session, err := mgo.Dial(address)
	session.SetMode(mgo.Monotonic, true)
	return err, session.DB(database)
}
