package echo_repoimpl

import (
	"context"

	"github.com/furu2revival/musicbox/app/adapter/dao"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
	"github.com/furu2revival/musicbox/app/infrastructure/trace"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct{}

func NewRepository() repository.EchoRepository {
	return &Repository{}
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, echos ...model.Echo) error {
	ctx, span := trace.StartSpan(ctx, "echo_repoimpl.Save")
	defer span.End()

	dtos := make([]*dao.Echo, len(echos))
	for i, echo := range echos {
		dtos[i] = &dao.Echo{
			ID:        echo.ID.String(),
			Message:   echo.Message,
			Timestamp: echo.Timestamp,
		}
	}
	_, err := dao.EchoSlice(dtos).UpsertAll(ctx, tx, true, dao.EchoPrimaryKeyColumns, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
