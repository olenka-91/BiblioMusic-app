package service

import (
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
)

type Song interface {
	Create(groupName string, songName string, s domain.Song) (int, error)
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
