// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"context"
	"github.com/furu2revival/musicbox/app/adapter/handler"
	"github.com/furu2revival/musicbox/app/adapter/handler/debug/echo_handler"
	"github.com/furu2revival/musicbox/app/adapter/handler/music_sheet_handler"
	"github.com/furu2revival/musicbox/app/adapter/repoimpl"
	"github.com/furu2revival/musicbox/app/adapter/repoimpl/echo_repoimpl"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/aop"
	"github.com/furu2revival/musicbox/app/infrastructure/db"
	"github.com/furu2revival/musicbox/app/usecase"
	"github.com/furu2revival/musicbox/app/usecase/echo_usecase"
	"github.com/google/wire"
	"net/http"
)

// Injectors from wire.go:

func InitializeAPIServerMux(ctx context.Context) (*http.ServeMux, error) {
	proxy := aop.NewProxy()
	musicSheetServiceHandler := music_sheet_handler.NewHandler(proxy)
	connection, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	echoRepository := echo_repoimpl.NewRepository()
	usecase := echo_usecase.NewUsecase(connection, echoRepository)
	echoServiceHandler := echo_handler.NewHandler(usecase, proxy)
	serveMux := handler.New(musicSheetServiceHandler, echoServiceHandler)
	return serveMux, nil
}

// wire.go:

var SuperSet = wire.NewSet(repoimpl.SuperSet, usecase.SuperSet, aop.NewProxy, db.NewConnection)