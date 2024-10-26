package model

import (
	"testing"

	"github.com/furu2revival/musicbox/testutils/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMusicSheet(t *testing.T) {
	type args struct {
		id    uuid.UUID
		title string
		notes []Note
	}
	tests := []struct {
		name    string
		args    args
		want    MusicSheet
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "楽譜を作成できる",
			args: args{
				id:    faker.UUIDv5("ms1"),
				title: "title",
				notes: []Note{
					NewNote(PitchC3, PitchD3, PitchE3),
				},
			},
			want: MusicSheet{
				ID:    faker.UUIDv5("ms1"),
				Title: "title",
				Notes: []Note{
					{
						PitchC3,
						PitchD3,
						PitchE3,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "len(title) < 1 の場合 => エラー",
			args: args{
				id:    uuid.New(),
				title: "",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrMusicTitleInvalid)
			},
		},
		{
			name: "len(title) > 100 の場合 => エラー",
			args: args{
				id: uuid.New(),
				title: func() string {
					s := ""
					for range 101 {
						s += "a"
					}
					return s
				}(),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrMusicTitleInvalid)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMusicSheet(tt.args.id, tt.args.title, tt.args.notes)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
