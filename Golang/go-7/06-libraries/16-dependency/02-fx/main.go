package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

func main() {
	type Username string
	UserModule := fx.Provide(func() Username {
		return "john"
	})

	var user Username
	app := fx.New(
		UserModule,
		fx.NopLogger,
		fx.Populate(&user),
	)
	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}
	defer app.Stop(context.Background())

	fmt.Println(user)
}
