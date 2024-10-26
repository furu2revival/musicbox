package music_sheet_usecase

import (
	"context"

	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
)

func (u Usecase) Create(ctx context.Context, rctx request_context.RequestContext, title string, notes []model.Note) (model.MusicSheet, error) {
	musicSheet, err := model.NewMusicSheet(rctx.IdempotencyKey(), title, notes)
	if err != nil {
		return model.MusicSheet{}, err
	}
	err = u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		err = u.musicSheetRepo.Save(ctx, tx, musicSheet)
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
