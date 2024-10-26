package faker

import (
	"time"

	"github.com/furu2revival/musicbox/app/core/request_context"
	"github.com/google/uuid"
)

type RequestContextBuilder struct {
	data request_context.RequestContext
}

func NewRequestContextBuilder() *RequestContextBuilder {
	return &RequestContextBuilder{
		data: request_context.NewRequestContext(uuid.New(), time.Now()),
	}
}

func (b *RequestContextBuilder) Build() request_context.RequestContext {
	return b.data
}

func (b *RequestContextBuilder) IdempotencyKey(idempotencyKey uuid.UUID) *RequestContextBuilder {
	b.data = request_context.NewRequestContext(idempotencyKey, b.data.Now())
	return b
}

func (b *RequestContextBuilder) Now(now time.Time) *RequestContextBuilder {
	b.data = request_context.NewRequestContext(b.data.IdempotencyKey(), now)
	return b
}
