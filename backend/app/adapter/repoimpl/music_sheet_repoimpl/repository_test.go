package music_sheet_repoimpl_test

import (
	"context"
	"testing"

	"github.com/furu2revival/musicbox/app/adapter/dao"
	"github.com/furu2revival/musicbox/app/adapter/repoimpl/music_sheet_repoimpl"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
	"github.com/furu2revival/musicbox/testutils"
	"github.com/furu2revival/musicbox/testutils/bdd"
	"github.com/furu2revival/musicbox/testutils/faker"
	"github.com/furu2revival/musicbox/testutils/fixture"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func TestRepository_Get(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type given struct {
		seeds []fixture.Seed
	}
	type when struct {
		id uuid.UUID
	}
	type then = func(t *testing.T, got model.MusicSheet, err error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "レコードが存在する状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.MusicSheet{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Title:        "ms1",
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
						id: faker.UUIDv5("ms1"),
					},
					Then: func(t *testing.T, got model.MusicSheet, err error) {
						require.NoError(t, err)

						want := model.MusicSheet{
							ID:    faker.UUIDv5("ms1"),
							Title: "ms1",
							Notes: []model.Note{
								{
									model.PitchC3,
								},
								{
									model.PitchC3,
									model.PitchD3,
								},
							},
						}
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "ID が存在しない => エラー",
					When: when{
						id: faker.UUIDv5("not-exist"),
					},
					Then: func(t *testing.T, got model.MusicSheet, err error) {
						require.ErrorIs(t, err, repository.ErrMusicSheetNotFound)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), given.seeds...)

			var got model.MusicSheet
			err := conn.BeginRoTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := music_sheet_repoimpl.NewRepository()

				var err error
				got, err = r.Get(ctx, tx, when.id)
				if err != nil {
					return err
				}
				return nil
			})
			then(t, got, err)
		})
	}
}

func TestRepository_Save(t *testing.T) {
	conn := testutils.MustDBConn(t)

	type given struct {
		seeds []fixture.Seed
	}
	type when struct {
		musicSheet model.MusicSheet
	}
	type then = func(t *testing.T, got model.MusicSheet, err error)
	tests := []bdd.Testcase[given, when, then]{
		{
			Name: "レコードが存在する状態で",
			Given: given{
				seeds: []fixture.Seed{
					&dao.MusicSheet{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Title:        "ms1",
					},
					&dao.Note{
						Index:        0,
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Pitches: types.Int64Array{
							int64(model.PitchC3),
						},
					},
				},
			},
			Behaviors: []bdd.Behavior[when, then]{
				{
					Name: "PK が重複する場合 => 更新する",
					When: when{
						musicSheet: model.MusicSheet{
							ID:    faker.UUIDv5("ms1"),
							Title: "ms1",
							Notes: []model.Note{
								{
									model.PitchA4,
									model.PitchB4,
								},
								{
									model.PitchC4,
								},
								{
									model.PitchC3,
								},
							},
						},
					},
					Then: func(t *testing.T, got model.MusicSheet, err error) {
						require.NoError(t, err)

						want := model.MusicSheet{
							ID:    faker.UUIDv5("ms1"),
							Title: "ms1",
							Notes: []model.Note{
								{
									model.PitchA4,
									model.PitchB4,
								},
								{
									model.PitchC4,
								},
								{
									model.PitchC3,
								},
							},
						}
						assert.Equal(t, want, got)
					},
				},
				{
					Name: "PK が重複しない場合 => 新規作成する",
					When: when{
						musicSheet: model.MusicSheet{
							ID:    faker.UUIDv5("ms2"),
							Title: "ms2",
							Notes: []model.Note{
								{
									model.PitchA4,
									model.PitchB4,
								},
								{
									model.PitchC4,
								},
								{
									model.PitchC3,
								},
							},
						},
					},
					Then: func(t *testing.T, got model.MusicSheet, err error) {
						require.NoError(t, err)

						want := model.MusicSheet{
							ID:    faker.UUIDv5("ms2"),
							Title: "ms2",
							Notes: []model.Note{
								{
									model.PitchA4,
									model.PitchB4,
								},
								{
									model.PitchC4,
								},
								{
									model.PitchC3,
								},
							},
						}
						assert.Equal(t, want, got)
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run(t, func(t *testing.T, given given, when when, then then) {
			defer testutils.Teardown(t)
			fixture.SetupSeeds(t, context.Background(), given.seeds...)

			var got model.MusicSheet
			err := conn.BeginRwTransaction(context.Background(), func(ctx context.Context, tx transaction.Transaction) error {
				r := music_sheet_repoimpl.NewRepository()
				err := r.Save(ctx, tx, when.musicSheet)
				if err != nil {
					return err
				}

				got, err = r.Get(ctx, tx, when.musicSheet.ID)
				if err != nil {
					return err
				}
				return nil
			})
			then(t, got, err)
		})
	}
}

func Test_CheckNoteDiff(t *testing.T) {
	type args struct {
		newDtos     []*dao.Note
		currentDtos []*dao.Note
	}
	tests := []struct {
		name         string
		args         args
		wantUpserted dao.NoteSlice
		wantDeleted  dao.NoteSlice
	}{
		{
			name: "作成/更新/削除された音符が存在する => upserted と deleted が正しく計算される",
			args: args{
				// index1: 更新
				// index2: 削除
				// index3: 作成
				newDtos: []*dao.Note{
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        1,
						Pitches:      types.Int64Array{int64(model.PitchC3)},
					},
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        3,
						Pitches:      types.Int64Array{int64(model.PitchC3)},
					},
				},
				currentDtos: []*dao.Note{
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        1,
						Pitches:      types.Int64Array{int64(model.PitchC4)},
					},
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        2,
						Pitches:      types.Int64Array{int64(model.PitchC3)},
					},
				},
			},
			wantUpserted: []*dao.Note{
				{
					MusicSheetID: faker.UUIDv5("ms1").String(),
					Index:        1,
					Pitches:      types.Int64Array{int64(model.PitchC3)},
				},
				{
					MusicSheetID: faker.UUIDv5("ms1").String(),
					Index:        3,
					Pitches:      types.Int64Array{int64(model.PitchC3)},
				},
			},
			wantDeleted: []*dao.Note{
				{
					MusicSheetID: faker.UUIDv5("ms1").String(),
					Index:        2,
					Pitches:      types.Int64Array{int64(model.PitchC3)},
				},
			},
		},
		{
			name: "全ての音符が完全一致 => 何もしない",
			args: args{
				newDtos: []*dao.Note{
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        1,
					},
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        2,
					},
				},
				currentDtos: []*dao.Note{
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        1,
					},
					{
						MusicSheetID: faker.UUIDv5("ms1").String(),
						Index:        2,
					},
				},
			},
			wantUpserted: []*dao.Note{},
			wantDeleted:  []*dao.Note{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUpserted, gotDeleted := music_sheet_repoimpl.CheckNoteDiff(tt.args.newDtos, tt.args.currentDtos)
			assert.Equal(t, tt.wantUpserted, gotUpserted)
			assert.Equal(t, tt.wantDeleted, gotDeleted)
		})
	}
}
