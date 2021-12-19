/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:53 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package container

import (
	"github.com/robiokidenis/microservice-mvc-1/httphandler/auth"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/extra"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/role"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/roleuser"
	"github.com/robiokidenis/microservice-mvc-1/httphandler/user"
	"go.uber.org/dig"
)

func BuildHttpHandlerProvider(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(extra.NewHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(auth.NewHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(user.NewHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(role.NewHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(roleuser.NewHandler); err != nil {
		panic(err)
	}

	return container
}
