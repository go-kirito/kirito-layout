package di

import (
	"github.com/go-kirito/pkg/application"
	"github.com/google/wire"
)

type UseCases struct {
}

func RegisterService(app application.Application) error {
	_, err := MakeUseCase()
	if err != nil {
		return err
	}

	return nil
}

func MakeUseCase() (*UseCases, error) {
	panic(wire.Build(
		wire.Struct(new(UseCases), "*"),
	))
}
