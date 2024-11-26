package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
)

type Song interface {
	Create(s domain.Song) (int, error)
	GetSongsList(s domain.PaginatedSongInput) ([]domain.Song, error)
	GetSongText(s domain.PaginatedSongTextInput) (domain.PaginatedSongTextResponse, error)
	Delete(songID int) error
	Update(songID int, input domain.SongUpdateInput) error
}

type Repository struct {
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Song: NewSongPostgres(db)}
}
