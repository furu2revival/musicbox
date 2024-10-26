package repoimpl

import (
	"github.com/furu2revival/musicbox/app/adapter/repoimpl/echo_repoimpl"
	"github.com/furu2revival/musicbox/app/adapter/repoimpl/music_sheet_repoimpl"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	echo_repoimpl.NewRepository,
	music_sheet_repoimpl.NewRepository,
)
