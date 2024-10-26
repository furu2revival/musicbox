package music_sheet_handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/furu2revival/musicbox/app/adapter/dao"
	"github.com/furu2revival/musicbox/app/registry"
	"github.com/furu2revival/musicbox/protobuf/api"
	"github.com/furu2revival/musicbox/protobuf/api/api_errors"
	"github.com/furu2revival/musicbox/protobuf/api/apiconnect"
	"github.com/furu2revival/musicbox/testutils"
	"github.com/furu2revival/musicbox/testutils/bdd"
	"github.com/furu2revival/musicbox/testutils/faker"
	"github.com/furu2revival/musicbox/testutils/fixture"
	"github.com/furu2revival/musicbox/testutils/testconnect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handler_CreateV1(t *testing.T) {
	mux, err := registry.InitializeAPIServerMux(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(mux)
	defer server.Close()

	type given struct {
		seeds []fixture.Seed
	}
	type when struct {
		req  *api.MusicSheetServiceCreateV1Request
		opts []testconnect.Option
	}
	type then = func(*testing.T, *connect.Response[api.MusicSheetServiceCreateV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "楽譜が存在する状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.MusicSheet{
						MusicSheetID: faker.UUIDv5("ms1").String(),
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "冪等キーと同じ楽譜IDが存在しない場合 => 新規作成できる",
					When: when{
						req: &api.MusicSheetServiceCreateV1Request{
							Title: "test",
							Notes: []*api.Note{},
						},
						opts: []testconnect.Option{
							testconnect.WithIdempotencyKey(faker.UUIDv5("ms2")),
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.MusicSheetServiceCreateV1Response], err error) {
						require.NoError(t, err)

						want := &api.MusicSheetServiceCreateV1Response{
							MusicSheetId: faker.UUIDv5("ms2").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "冪等キーと同じ楽譜IDが存在する場合 => 冪等に処理する",
					When: when{
						req: &api.MusicSheetServiceCreateV1Request{
							Title: "test",
							Notes: []*api.Note{},
						},
						opts: []testconnect.Option{
							testconnect.WithIdempotencyKey(faker.UUIDv5("ms1")),
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.MusicSheetServiceCreateV1Response], err error) {
						require.NoError(t, err)

						want := &api.MusicSheetServiceCreateV1Response{
							MusicSheetId: faker.UUIDv5("ms1").String(),
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "不正なタイトル => エラー",
					When: when{
						req: &api.MusicSheetServiceCreateV1Request{
							Title: "",
							Notes: []*api.Note{},
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.MusicSheetServiceCreateV1Response], err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_METHOD_ILLEGAL_ARGUMENT, err)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			fixture.SetupSeeds(t, context.Background(), given.seeds...)
			defer testutils.Teardown(t)

			got, err := testconnect.MethodInvoke(
				apiconnect.NewMusicSheetServiceClient(http.DefaultClient, server.URL).CreateV1,
				when.req,
				when.opts...,
			)
			then(t, got, err)
		})
	}
}
