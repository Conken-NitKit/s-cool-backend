package di

import (
	"condog/internal/config"
	"context"
	"errors"
)

type App interface {
	Run() error
}

type app struct {
	ctx        context.Context
	cfg        config.AppConfig
}

func ProvideApp(
	ctx context.Context,
	cfg config.AppConfig,
) App {
	return &app{
		ctx: ctx,
		cfg: cfg,
	}
}

func (a *app) Run() error {
	return errors.New("サンプルエラー")
}
