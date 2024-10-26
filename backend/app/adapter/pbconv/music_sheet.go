package pbconv

import (
	"github.com/furu2revival/musicbox/app/domain/model"
	"github.com/furu2revival/musicbox/protobuf/api"
)

func FromNotePbs(v []*api.Note) []model.Note {
	res := make([]model.Note, len(v))
	for i, note := range v {
		note.GetPitches()
		res[i] = FromPitchPbs(note.GetPitches())
	}
	return res
}

func FromPitchPbs(v []api.Pitch) []model.Pitch {
	res := make([]model.Pitch, 0)
	for _, pitch := range v {
		switch pitch {
		case api.Pitch_PITCH_C3:
			res = append(res, model.PitchC3)
		case api.Pitch_PITCH_D3:
			res = append(res, model.PitchD3)
		case api.Pitch_PITCH_E3:
			res = append(res, model.PitchE3)
		case api.Pitch_PITCH_F3:
			res = append(res, model.PitchF3)
		case api.Pitch_PITCH_G3:
			res = append(res, model.PitchG3)
		case api.Pitch_PITCH_A4:
			res = append(res, model.PitchA4)
		case api.Pitch_PITCH_B4:
			res = append(res, model.PitchB4)
		case api.Pitch_PITCH_C4:
			res = append(res, model.PitchC4)
		case api.Pitch_PITCH_UNSPECIFIED:
			// do nothing
		}
	}
	return res
}

func ToMusicSheetPb(v model.MusicSheet) *api.MusicSheet {
	return &api.MusicSheet{
		MusicSheetId: v.ID.String(),
		Title:        v.Title,
		Notes:        ToNotePbs(v.Notes),
	}
}

func ToNotePbs(v []model.Note) []*api.Note {
	res := make([]*api.Note, len(v))
	for i, note := range v {
		res[i] = &api.Note{
			Pitches: ToPitchPbs(note),
		}
	}
	return res
}

func ToPitchPbs(v []model.Pitch) []api.Pitch {
	res := make([]api.Pitch, len(v))
	for i, pitch := range v {
		switch pitch {
		case model.PitchC3:
			res[i] = api.Pitch_PITCH_C3
		case model.PitchD3:
			res[i] = api.Pitch_PITCH_D3
		case model.PitchE3:
			res[i] = api.Pitch_PITCH_E3
		case model.PitchF3:
			res[i] = api.Pitch_PITCH_F3
		case model.PitchG3:
			res[i] = api.Pitch_PITCH_G3
		case model.PitchA4:
			res[i] = api.Pitch_PITCH_A4
		case model.PitchB4:
			res[i] = api.Pitch_PITCH_B4
		case model.PitchC4:
			res[i] = api.Pitch_PITCH_C4
		}
	}
	return res
}
