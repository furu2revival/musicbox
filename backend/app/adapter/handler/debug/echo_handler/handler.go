package echo_handler

import (
	"context"

	"github.com/furu2revival/musicbox/app/infrastructure/connect/aop"
	"github.com/furu2revival/musicbox/app/usecase/echo_usecase"
	"github.com/furu2revival/musicbox/protobuf/api/debug"
	"github.com/furu2revival/musicbox/protobuf/api/debug/debugconnect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handler struct {
	uc *echo_usecase.Usecase
}

func NewHandler(uc *echo_usecase.Usecase, proxy aop.Proxy) debugconnect.EchoServiceHandler {
	return debug.NewEchoServiceHandler(&handler{uc: uc}, proxy)
}

func (h handler) EchoV1(ctx context.Context, req *aop.Request[*debug.EchoServiceEchoV1Request]) (*debug.EchoServiceEchoV1Response, error) {
	result, err := h.uc.Echo(ctx, req.RequestContext(), req.Msg().GetMessage())
	if err != nil {
		return nil, err
	}
	return &debug.EchoServiceEchoV1Response{
		Message:   result.Message,
		Timestamp: timestamppb.New(result.Timestamp),
	}, nil
}
