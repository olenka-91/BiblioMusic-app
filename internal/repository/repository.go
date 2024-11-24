package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
)

type Song interface {
	Create(s domain.SongList) (int, error)
	GetSongsList(s domain.PaginatedSongInput) ([]domain.SongOutput, error)
	GetSongText(s domain.PaginatedSongTextInput) (domain.PaginatedSongTextResponse, error)
	//GetByID(userID int, remindID int) (domain.Remind, error)
	//GetAll(userID int) ([]domain.Remind, error)
	//Delete(userID, remindID int) error
	//Update(userID, remindID int, input domain.RemindUpdateInput) error
}

type Repository struct {
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Song: NewSongPostgres(db)}
}
