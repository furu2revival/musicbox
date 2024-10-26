package music_sheet_usecase

import (
	"context"
	
	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
)

func (u Usecase) Get(ctx context.Context, rctx request_context.RequestContext) (model.MusicSheet, error) {
	
}
