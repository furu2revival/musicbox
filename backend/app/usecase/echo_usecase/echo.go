package echo_usecase

import (
	"context"
	"fmt"

	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
)

func (u Usecase) Echo(ctx context.Context, rctx request_context.RequestContext, message string) (model.Echo, error) {
	echo := model.NewEcho(message, rctx.Now())
	err := u.conn.BeginRwTransaction(ctx, func(ctx context.Context, tx transaction.Transaction) error {
		err := u.echoRepo.Save(ctx, tx, echo)
		if err != nil {
			return fmt.Errorf("echoRepo.Save failed, %w", err)
		}
		return nil
	})
	if err != nil {
		return model.Echo{}, err
	}
	return echo, nil
}
