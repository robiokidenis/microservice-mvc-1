/*
 * Copyright (c) 2021
 * Created at 5/20/21 10:31 PM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package container

import (
	"github.com/robiokidenis/microservice-mvc-1/domain/services"
	"go.uber.org/dig"
)

func BuildRepositoryProvider(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(services.NewUserServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewUserSocialServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewRoleServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewRoleUserService); err != nil {
		panic(err)
	}

	return container
}
