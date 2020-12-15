package app

import (
	"github.com/wajox/clonegopkg/internal/app/dependencies"
	"github.com/wajox/clonegopkg/internal/app/initializers"
)

type Application struct {
	Container *dependencies.Container
}

func InitializeApplication() (*Application, error) {
	if err := initializers.InitializeEnvs(); err != nil {
		return nil, err
	}

	if err := initializers.InitializeLogs(); err != nil {
		return nil, err
	}

	return BuildApplication()
}
