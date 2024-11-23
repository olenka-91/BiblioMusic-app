package service

import (
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
)

type Song interface {
	Create(s domain.SongList) (int, error)
	GetSongsList(domain.PaginatedSongInput) ([]domain.SongOutput, error)
	//GetByID(userID int, remindID int) (domain.Remind, error)
	//GetAll(userID int) ([]domain.Remind, error)
	//Delete(userID, remindID int) error
	//Update(userID, remindID int, input domain.RemindUpdateInput) error
}

type Service struct {
	Song
}

func NewService(r *repository.Repository) *Service {
	return &Service{Song: NewSongService(r.Song)}
}
