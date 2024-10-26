package repository

import (
	"context"

	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
)

type EchoRepository interface {
	Save(ctx context.Context, tx transaction.Transaction, echos ...model.Echo) error
}
