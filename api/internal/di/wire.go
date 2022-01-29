//go:build wireinject
// +build wireinject

package di

import (
	"condog/internal/config"
	"context"

	"github.com/google/wire"
)

func InitializeApp(ctx context.Context, cfg config.AppConfig) (App, error) {
	wire.Build(
		ProvideApp,
	)
	return nil, nil
}
