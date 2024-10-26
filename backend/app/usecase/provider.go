package usecase

import (
	"github.com/furu2revival/musicbox/app/usecase/echo_usecase"
	"github.com/furu2revival/musicbox/app/usecase/music_sheet_usecase"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	echo_usecase.NewUsecase,
	music_sheet_usecase.NewUsecase,
)
