// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"binanceexchange_user/internal/biz"
	"binanceexchange_user/internal/conf"
	"binanceexchange_user/internal/data"
	"binanceexchange_user/internal/server"
	"binanceexchange_user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	binanceUserRepo := data.NewBinanceUserRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	binanceUserUsecase := biz.NewBinanceDataUsecase(binanceUserRepo, transaction, logger)
	binanceUserService := service.NewBinanceDataService(binanceUserUsecase)
	httpServer := server.NewHTTPServer(confServer, binanceUserService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
