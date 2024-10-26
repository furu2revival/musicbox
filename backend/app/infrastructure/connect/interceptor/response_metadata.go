package interceptor

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/mdval"
)

func NewResponseMetadataInterceptor(serverVersion string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			resp, err := next(ctx, req)
			if err != nil {
				return nil, err
			}

			// TODO: リクエストIDをレスポンスヘッダに追加する
			mdval.SetOutgoingHeader(resp, mdval.OutgoingHeaderMD{
				mdval.RespondTimestampKey: time.Now().UTC().Format(time.RFC3339Nano),
				mdval.ServerVersionKey:    serverVersion,
			})
			return resp, nil
		}
	}
}
