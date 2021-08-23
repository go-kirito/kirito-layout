// +build wireinject

/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/19 16:08
 */
package di

import (
	"app/api/helloworld"
	"app/internal/app/helloworld/usecase"

	"github.com/go-kirito/pkg/application"
	"github.com/google/wire"
)

func RegisterService(app application.Application) error {
	wire.Build(helloworld.RegisterGreeterServer, usecase.NewHelloWorldUseCase)
	return nil
}
