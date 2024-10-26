package handler

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/furu2revival/musicbox/app/adapter/handler/debug/echo_handler"
	"github.com/furu2revival/musicbox/app/adapter/handler/music_sheet_handler"
	"github.com/furu2revival/musicbox/app/core/config"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/interceptor"
	"github.com/furu2revival/musicbox/protobuf/api/apiconnect"
	"github.com/furu2revival/musicbox/protobuf/api/debug/debugconnect"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	music_sheet_handler.NewHandler,
	echo_handler.NewHandler,
	New,
)

func New(
	musicSheetHandler apiconnect.MusicSheetServiceHandler,
	echoHandler debugconnect.EchoServiceHandler,
) *http.ServeMux {
	opts := connect.WithInterceptors(interceptor.New()...)
	mux := http.NewServeMux()
	mux.Handle(apiconnect.NewMusicSheetServiceHandler(musicSheetHandler, opts))
	if config.Get().GetDebug() {
		mux.Handle(debugconnect.NewEchoServiceHandler(echoHandler, opts))
	}
	return mux
}
