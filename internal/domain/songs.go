package domain

import (
	"errors"
)

type Song struct {
	ID          int64  `json:"id"`
	GroupID     int64  `json:"group_id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
}

type SongUpdateInput struct {
	Title       *string `json:"title"`
	Text        *string `json:"text"`
	ReleaseDate *string `json:"release_date"`
	Link        *string `json:"link"`
}

func (r *SongUpdateInput) Validate() error {
	if (r.Title == nil) && (r.Text == nil) && (r.ReleaseDate == nil) && (r.Link == nil) {
		return errors.New("update structure is empty")
	}
	return nil
}
