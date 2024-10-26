package music_sheet_handler

import (
	"context"

	"github.com/furu2revival/musicbox/app/infrastructure/connect/aop"
	"github.com/furu2revival/musicbox/protobuf/api"
	"github.com/furu2revival/musicbox/protobuf/api/apiconnect"
)

type handler struct {
}

func NewHandler(proxy aop.Proxy) apiconnect.MusicSheetServiceHandler {
	return api.NewMusicSheetServiceHandler(&handler{}, proxy)
}

func (h handler) GetV1(ctx context.Context, req *aop.Request[*api.MusicSheetServiceGetV1Request]) (*api.MusicSheetServiceGetV1Response, error) {
	// TODO: implement me
	return &api.MusicSheetServiceGetV1Response{}, nil
}

func (h handler) GetV1Errors(errs *api.MusicSheetServiceGetV1Errors) {
	// TODO: implement me
}

func (h handler) CreateV1(ctx context.Context, req *aop.Request[*api.MusicSheetServiceCreateV1Request]) (*api.MusicSheetServiceCreateV1Response, error) {
	// TODO: implement me
	return &api.MusicSheetServiceCreateV1Response{}, nil
}

func (h handler) CreateV1Errors(errs *api.MusicSheetServiceCreateV1Errors) {
	// TODO: implement me
}
