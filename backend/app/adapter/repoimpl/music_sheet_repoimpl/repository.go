package music_sheet_repoimpl

import (
	"context"

	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
	"github.com/google/uuid"
)

type Repository struct{}

func NewRepository() repository.MusicSheetRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, id uuid.UUID) (model.MusicSheet, error) {
	// TODO: #12 MusicSheetRepository.Get() を実装する
	return model.MusicSheet{}, repository.ErrMusicSheetNotFound
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, musicSheet model.MusicSheet) error {
	// TODO: #13 MusicSheetRepository.Save() を実装する
	return nil
}
