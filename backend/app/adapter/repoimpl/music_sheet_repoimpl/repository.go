package music_sheet_repoimpl

import (
	"context"
	"database/sql"
	"errors"

	"github.com/furu2revival/musicbox/app/adapter/dao"
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/app/domain/repository"
	"github.com/furu2revival/musicbox/app/domain/repository/transaction"
	"github.com/furu2revival/musicbox/pkg/vector"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type Repository struct{}

func NewRepository() repository.MusicSheetRepository {
	return &Repository{}
}

func (r Repository) Get(ctx context.Context, tx transaction.Transaction, id uuid.UUID) (model.MusicSheet, error) {
	musicSheet, err := dao.MusicSheets(dao.MusicSheetWhere.MusicSheetID.EQ(id.String())).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.MusicSheet{}, repository.ErrMusicSheetNotFound
		}
		return model.MusicSheet{}, err
	}
	noteDtos, err := dao.Notes(dao.NoteWhere.MusicSheetID.EQ(musicSheet.MusicSheetID), qm.OrderBy("Index")).All(ctx, tx)
	if err != nil {
		return model.MusicSheet{}, err
	}
	notes := make([]model.Note, musicSheet.NumberOfNotes)
	for _, note := range noteDtos {
		notes[note.Index] = model.NewNote(vector.Map(note.Pitches, func(pitch int64) model.Pitch {
			return model.Pitch(pitch)
		})...)
	}
	return model.NewMusicSheet(uuid.MustParse(musicSheet.MusicSheetID), musicSheet.Title, notes)
}

func (r Repository) Save(ctx context.Context, tx transaction.Transaction, musicSheet model.MusicSheet) error {
	{
		dto := &dao.MusicSheet{
			MusicSheetID:  musicSheet.ID.String(),
			Title:         musicSheet.Title,
			NumberOfNotes: len(musicSheet.Notes),
		}
		err := dto.Upsert(ctx, tx, true, dao.MusicSheetPrimaryKeyColumns, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
	}
	{
		currentDtos, err := dao.Notes(dao.NoteWhere.MusicSheetID.EQ(musicSheet.ID.String())).All(ctx, tx)
		if err != nil {
			return err
		}
		newDtos := make([]*dao.Note, 0)
		for i, note := range musicSheet.Notes {
			pitches := make(types.Int64Array, len(note))
			for j, pitch := range note {
				pitches[j] = int64(pitch)
			}
			newDtos = append(newDtos, &dao.Note{
				Index:        i,
				MusicSheetID: musicSheet.ID.String(),
				Pitches:      pitches,
			})
		}

		upserted, deleted := CheckNoteDiff(newDtos, currentDtos)
		_, err = upserted.UpsertAll(ctx, tx, true, dao.NotePrimaryKeyColumns, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}

		_, err = deleted.DeleteAll(ctx, tx)
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckNoteDiff は、新しい音符と現在の音符を比較して、作成/更新/削除された音符を仕分けます。
func CheckNoteDiff(newDtos, currentDtos dao.NoteSlice) (upserted dao.NoteSlice, deleted dao.NoteSlice) {
	currentMap := make(map[int]*dao.Note)
	for _, current := range currentDtos {
		currentMap[current.Index] = current
	}

	upserted = make([]*dao.Note, 0)
	for _, dto := range newDtos {
		current, ok := currentMap[dto.Index]
		if !ok {
			// 作成された
			upserted = append(upserted, dto)
		} else if !cmp.Equal(dto, current, cmpopts.IgnoreFields(dao.Note{}, "CreatedAt", "UpdatedAt")) {
			// 更新された
			upserted = append(upserted, dto)
		}
		delete(currentMap, dto.Index)
	}

	deleted = make([]*dao.Note, 0)
	for _, leftover := range currentMap {
		deleted = append(deleted, leftover)
	}
	return upserted, deleted
}
