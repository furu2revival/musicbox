package music_sheet_usecase

import (
    "context"

    "github.com/furu2revival/musicbox/app/core/request_context"
    "github.com/furu2revival/musicbox/app/domain/model"
    "github.com/furu2revival/musicbox/app/adapter/repoimpl/music_sheet_repoimpl"
	"github.com/google/uuid"
)

func (u Usecase) Get(ctx context.Context, rctx request_context.RequestContext, id uuid.UUID) (model.MusicSheet, error) {
	
}