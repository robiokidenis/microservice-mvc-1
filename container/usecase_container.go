/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:36 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package container

import (
	"github.com/robiokidenis/microservice-mvc-1/app/context/auth"
	"github.com/robiokidenis/microservice-mvc-1/app/context/role"
	"github.com/robiokidenis/microservice-mvc-1/app/context/roleuser"
	"github.com/robiokidenis/microservice-mvc-1/app/context/user"
	"go.uber.org/dig"
)

func BuildUsecaseProvider(container *dig.Container) *dig.Container {
	/** register your use case here with this format */
	var err error

	if err = container.Provide(auth.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(user.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(role.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(roleuser.NewUsecase); err != nil {
		panic(err)
	}
	return container
}
