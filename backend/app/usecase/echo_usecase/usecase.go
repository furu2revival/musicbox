package echo_usecase

import (
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
)

type Usecase struct {
	conn     transaction.Connection
	echoRepo repository.EchoRepository
}

func NewUsecase(conn transaction.Connection, echoRepo repository.EchoRepository) *Usecase {
	return &Usecase{
		conn:     conn,
		echoRepo: echoRepo,
	}
}
