//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/wajox/clonegopkg/internal/app/dependencies"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		wire.Struct(new(dependencies.Container)),
		wire.Struct(new(Application), "Container"),
	)

	return &Application{}, nil
}
