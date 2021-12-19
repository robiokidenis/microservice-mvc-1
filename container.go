/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:02 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package main

import (
	"github.com/robiokidenis/microservice-mvc-1/container"
	"go.uber.org/dig"
)

func NewBuildContainer() *dig.Container {
	c := dig.New()
	c = container.BuildConfigProvider(c)
	c = container.BuildRepositoryProvider(c)
	c = container.BuildUsecaseProvider(c)
	c = container.BuildHttpHandlerProvider(c)

	return c
}
