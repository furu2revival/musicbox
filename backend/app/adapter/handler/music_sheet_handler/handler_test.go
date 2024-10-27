package music_sheet_handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/furu2revival/musicbox/app/adapter/dao"
	"github.com/furu2revival/musicbox/app/domain/model"
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
	"github.com/volatiletech/sqlboiler/v4/types"
)

func Test_handler_GetV1(t *testing.T) {
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
		req  *api.MusicSheetServiceGetV1Request
		opts []testconnect.Option
	}
	type then = func(*testing.T, *connect.Response[api.MusicSheetServiceGetV1Response], error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "楽譜が存在する状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.MusicSheet{
						MusicSheetID:  faker.UUIDv5("ms1").String(),
						Title:         "ms1",
						NumberOfNotes: 2,
					},
					&dao.Note{
						Index:        0,
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Pitches: types.Int64Array{
							int64(model.PitchC3),
						},
					},
					&dao.Note{
						Index:        1,
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Pitches: types.Int64Array{
							int64(model.PitchC3),
							int64(model.PitchD3),
						},
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "ID が存在する => 取得できる",
					When: when{
						req: &api.MusicSheetServiceGetV1Request{
							MusicSheetId: faker.UUIDv5("ms1").String(),
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.MusicSheetServiceGetV1Response], err error) {
						require.NoError(t, err)

						want := &api.MusicSheetServiceGetV1Response{
							MusicSheet: &api.MusicSheet{
								MusicSheetId: faker.UUIDv5("ms1").String(),
								Title:        "ms1",
								Notes: []*api.Note{
									{
										Pitches: []api.Pitch{api.Pitch_PITCH_C3},
									},
									{
										Pitches: []api.Pitch{api.Pitch_PITCH_C3, api.Pitch_PITCH_D3},
									},
								},
							},
						}
						assert.EqualExportedValues(t, want, got.Msg)
					},
				},
				{
					Name: "ID が存在しない => エラー",
					When: when{
						req: &api.MusicSheetServiceGetV1Request{
							MusicSheetId: faker.UUIDv5("not-exists").String(),
						},
					},
					Then: func(t *testing.T, got *connect.Response[api.MusicSheetServiceGetV1Response], err error) {
						testconnect.AssertErrorCode(t, api_errors.ErrorCode_METHOD_RESOURCE_NOT_FOUND, err)
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
				apiconnect.NewMusicSheetServiceClient(http.DefaultClient, server.URL).GetV1,
				when.req,
				when.opts...,
			)
			then(t, got, err)
		})
	}
}

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
