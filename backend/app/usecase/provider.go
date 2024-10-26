package usecase

import (
	"github.com/furu2revival/musicbox/app/usecase/echo_usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	echo_usecase.NewUsecase,
)
