package domain

import (
	"errors"
)

type Song struct {
	ID           int64  `json:"id"`
	Group_ID     int64  `json:"group_id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	Release_Date string `json:"release_date"`
	Link         string `json:"link"`
}

type SongOutput struct {
	GroupName   *string `json:"group_name"`
	Title       *string `json:"title"`
	Text        *string `json:"text"`
	ReleaseDate *string `json:"release_date"`
	Link        *string `json:"link"`
}

type SongUpdateInput struct {
	Title       *string `json:"title"`
	Text        *string `json:"text"`
	ReleaseDate *string `json:"release_date"`
	Link        *string `json:"link"`
}

type SongDetail struct {
	Text        *string `json:"text"`
	ReleaseDate *string `json:"release_date"`
	Link        *string `json:"link"`
}

type SongList struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

func (r *SongUpdateInput) Validate() error {
	if (r.Title == nil) && (r.Text == nil) && (r.ReleaseDate == nil) && (r.Link == nil) {
		return errors.New("update structure is empty")
	}
	return nil
}

type PaginatedSongResponse struct {
	Data       []SongOutput `json:"data"`
	TotalCount int          `json:"total_count"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
}

type PaginatedSongInput struct {
	GroupName   string `json:"group_name"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
}

type PaginatedSongTextInput struct {
	SongId   int `json:"id"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginatedSongTextResponse struct {
	Title       string   `json:"title"`
	Verses      []string `json:"verses"`
	Page        int      `json:"page"`
	PageSize    int      `json:"page_size"`
	TotalVerses int      `json:"total_verses"`
}
