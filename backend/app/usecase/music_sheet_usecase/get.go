package music_sheet_usecase

import (
	"context"

	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
	"github.com/google/uuid"
)

func (u Usecase) Get(ctx context.Context, rctx request_context.RequestContext, id uuid.UUID) (model.MusicSheet, error) {
	var musicSheet model.MusicSheet
	err := u.conn.BeginRoTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		var err error
		musicSheet, err = u.musicSheetRepo.Get(ctx, tx, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.MusicSheet{}, err
	}
	return musicSheet, nil
}
