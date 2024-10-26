package connect

import (
	"context"
	"net/http"

	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/aop"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/mdval"
	"google.golang.org/protobuf/proto"
)

func Invoke[REQ, RES proto.Message](ctx context.Context, req REQ, header http.Header, info *aop.MethodInfo, method aop.Method[REQ, RES], proxy aop.Proxy) (RES, error) {
	var res RES
	wrap := func(ctx context.Context, rctx request_context.RequestContext, incomingMD mdval.IncomingMD) (proto.Message, error) {
		var err error
		res, err = method(ctx, aop.NewRequest(req, rctx))
		return res, err
	}
	return res, proxy(ctx, req, mdval.NewIncomingMD(header), info, wrap)
}
