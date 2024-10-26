package aop

import (
	"context"
	"errors"
	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/error_response"
	"github.com/furu2revival/musicbox/app/infrastructure/connect/mdval"
	"github.com/furu2revival/musicbox/protobuf/custom_option"
	"google.golang.org/protobuf/proto"
)

type (
	MethodOption                   = custom_option.MethodOption
	MethodErrDefinition            = custom_option.MethodErrorDefinition
	Method[REQ, RES proto.Message] func(context.Context, *Request[REQ]) (RES, error)

	// Proxy は、rpc method の前後で cross-cutting concern を実行するための関数です。
	// interceptor だと共通化しづらい処理を、ここで実行します。
	Proxy func(context.Context, proto.Message, mdval.IncomingMD, *MethodInfo, func(context.Context, request_context.RequestContext, mdval.IncomingMD) (proto.Message, error)) error
)

func NewProxy() Proxy {
	return func(ctx context.Context, req proto.Message, incomingMD mdval.IncomingMD, info *MethodInfo, method func(context.Context, request_context.RequestContext, mdval.IncomingMD) (proto.Message, error)) error {
		params, err := fixPreconditionParams(ctx, incomingMD)
		if err != nil {
			return err
		}
		rctx := params.RequestContext()

		_, err = method(ctx, rctx, incomingMD)
		if err != nil {
			if errDef, ok := info.FindErrorDefinition(err); ok {
				return error_response.New(errDef.GetCode(), errDef.GetSeverity(), errDef.GetMessage())
			}
			return err
		}
		return nil
	}
}

type MethodInfo struct {
	opt       *MethodOption
	errCauses map[error]*MethodErrDefinition
}

func NewMethodInfo(opt *MethodOption, errCauses map[error]*MethodErrDefinition) *MethodInfo {
	return &MethodInfo{
		opt:       opt,
		errCauses: errCauses,
	}
}

func (m *MethodInfo) Option() *MethodOption {
	return m.opt
}

func (m MethodInfo) FindErrorDefinition(err error) (*MethodErrDefinition, bool) {
	for cause, def := range m.errCauses {
		if errors.Is(err, cause) {
			return def, true
		}
	}
	return nil, false
}

type Request[T any] struct {
	msg  T
	rctx request_context.RequestContext
}

func NewRequest[T proto.Message](msg T, rctx request_context.RequestContext) *Request[T] {
	return &Request[T]{
		msg:  msg,
		rctx: rctx,
	}
}

func (r Request[T]) Msg() T {
	return r.msg
}

func (r Request[T]) RequestContext() request_context.RequestContext {
	return r.rctx
}
