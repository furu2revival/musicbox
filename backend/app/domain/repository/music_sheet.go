package repository

import (
	"context"
	"errors"

	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
	"github.com/google/uuid"
)

var (
	ErrMusicSheetNotFound = errors.New("music sheet not found")
)

type MusicSheetRepository interface {
	Get(ctx context.Context, tx transaction.Transaction, id uuid.UUID) (model.MusicSheet, error)
	Save(ctx context.Context, tx transaction.Transaction, musicSheet model.MusicSheet) error
}
