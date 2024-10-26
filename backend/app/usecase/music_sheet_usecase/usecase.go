package music_sheet_usecase

import (
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
)

type Usecase struct {
	conn           transaction.Connection
	musicSheetRepo repository.MusicSheetRepository
}

func NewUsecase(conn transaction.Connection, musicSheetRepo repository.MusicSheetRepository) *Usecase {
	return &Usecase{
		conn:           conn,
		musicSheetRepo: musicSheetRepo,
	}
}
