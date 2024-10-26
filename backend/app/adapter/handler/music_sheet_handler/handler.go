package music_sheet_handler

import (
	"context"

	"github.com/furu2revival/musicbox/app/adapter/pbconv"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/aop"
	"github.com/furu2revival/musicbox/app/usecase/music_sheet_usecase"
	"github.com/furu2revival/musicbox/protobuf/api"
	"github.com/furu2revival/musicbox/protobuf/api/apiconnect"
	"github.com/google/uuid"
)

type handler struct {
	uc *music_sheet_usecase.Usecase
}

func NewHandler(uc *music_sheet_usecase.Usecase, proxy aop.Proxy) apiconnect.MusicSheetServiceHandler {
	return api.NewMusicSheetServiceHandler(&handler{uc}, proxy)
}

func (h handler) GetV1(ctx context.Context, req *aop.Request[*api.MusicSheetServiceGetV1Request]) (*api.MusicSheetServiceGetV1Response, error) {
	result, err := h.uc.Get(ctx, req.RequestContext(), uuid.MustParse(req.Msg().GetMusicSheetId()))
	if err != nil {
		return nil, err
	}
	return &api.MusicSheetServiceGetV1Response{
		MusicSheet: pbconv.ToMusicSheetPb(result),
	}, nil
}

func (h handler) GetV1Errors(errs *api.MusicSheetServiceGetV1Errors) {
	errs.Map(repository.ErrMusicSheetNotFound, errs.RESOURCE_NOT_FOUND)
}

func (h handler) CreateV1(ctx context.Context, req *aop.Request[*api.MusicSheetServiceCreateV1Request]) (*api.MusicSheetServiceCreateV1Response, error) {
	result, err := h.uc.Create(ctx, req.RequestContext(), req.Msg().GetTitle(), pbconv.FromNotePbs(req.Msg().GetNotes()))
	if err != nil {
		return nil, err
	}
	return &api.MusicSheetServiceCreateV1Response{
		MusicSheetId: result.ID.String(),
	}, nil
}

func (h handler) CreateV1Errors(errs *api.MusicSheetServiceCreateV1Errors) {
	errs.Map(model.ErrMusicTitleInvalid, errs.ILLEGAL_ARGUMENT)
}
