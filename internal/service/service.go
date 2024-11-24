package service

import (
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
)

type Song interface {
	Create(s domain.SongList) (int, error)
	GetSongsList(domain.PaginatedSongInput) ([]domain.SongOutput, error)
	GetSongText(domain.PaginatedSongTextInput) (domain.PaginatedSongTextResponse, error)
	Delete(songID int) error
	Update(songID int, input domain.SongUpdateInput) error
}

type Service struct {
	Song
}

func NewService(r *repository.Repository) *Service {
	return &Service{Song: NewSongService(r.Song)}
}
