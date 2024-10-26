//go:build wireinject
// +build wireinject

package registry

import (
	"context"
	"net/http"

	"github.com/furu2revival/musicbox/app/adapter/handler"
	"github.com/furu2revival/musicbox/app/adapter/repoimpl"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/aop"
	"github.com/furu2revival/musicbox/app/infrastructure/db"
	"github.com/furu2revival/musicbox/app/usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	repoimpl.SuperSet,
	usecase.SuperSet,
	aop.NewProxy,
	db.NewConnection,
)

func InitializeAPIServerMux(ctx context.Context) (*http.ServeMux, error) {
	wire.Build(SuperSet, handler.SuperSet)
	return nil, nil
}
