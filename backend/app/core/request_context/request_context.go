package request_context

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var (
	_ json.Marshaler = RequestContext{}
)

// RequestContext は、機能によらずアプリケーション横断的なコンテキストを提供します。
type RequestContext struct {
	idempotencyKey uuid.UUID
	now            time.Time
}

func NewRequestContext(idempotencyKey uuid.UUID, now time.Time) RequestContext {
	return RequestContext{
		idempotencyKey: idempotencyKey,
		now:            now,
	}
}

// IdempotencyKey はトランザクションの冪等性を保証するために利用される、一意な識別子です。
// https://developer.mozilla.org/ja/docs/Glossary/Idempotent
func (c RequestContext) IdempotencyKey() uuid.UUID {
	return c.idempotencyKey
}

func (c RequestContext) Now() time.Time {
	return c.now
}

func (c RequestContext) JSON() map[string]interface{} {
	return map[string]interface{}{
		"idempotencyKey": c.idempotencyKey,
		"now":            c.now,
	}
}

func (c RequestContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		IdempotencyKey string `json:"idempotencyKey"`
		Now            string `json:"now"`
	}{
		IdempotencyKey: c.idempotencyKey.String(),
		Now:            c.now.Format(time.RFC3339Nano),
	})
}
