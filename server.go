/*
 * Copyright (c) 2021
 * Created at 5/20/21 11:18 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package main

import (
	"log"

	"github.com/robiokidenis/microservice-mvc-1/app/config"
	"github.com/gin-gonic/gin"
)

type ServerRoute struct {
	cfg    config.Config
	engine *gin.Engine
}

func NewRoute(cfg config.Config, engine *gin.Engine) *ServerRoute {
	return &ServerRoute{cfg, engine}
}

func (s *ServerRoute) Run() {
	if err := s.engine.Run(s.cfg.GetString(`server.address`)); err != nil {
		log.Fatal(err)
	}
}
