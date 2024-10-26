package model

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMusicTitleInvalid = errors.New("music title is invalid")
)

// MusicSheet は、楽譜を表します。
type MusicSheet struct {
	ID    uuid.UUID
	Title string
	Notes []Note
}

func NewMusicSheet(id uuid.UUID, title string, notes []Note) (MusicSheet, error) {
	if len(title) < 1 || len(title) > 100 {
		return MusicSheet{}, ErrMusicTitleInvalid
	}
	return MusicSheet{
		ID:    id,
		Title: title,
		Notes: notes,
	}, nil
}

// Note 音符
// ただし、ここでは 16 分音符のみを扱います。
type Note []Pitch

func NewNote(pitches ...Pitch) Note {
	return pitches
}

// Pitch 音階
type Pitch int

const (
	PitchC3 = iota + 1
	PitchD3
	PitchE3
	PitchF3
	PitchG3
	PitchA4
	PitchB4
	PitchC4
)
