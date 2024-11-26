package service

import (
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
)

type Song interface {
	Create(req domain.AddSongRequest, detail domain.SongDetail) (int, error)
	GetSongsList(s domain.PaginatedSongInput, page, pageSize string) ([]domain.Song, error)
	GetSongText(s domain.PaginatedSongTextInput, page, pageSize string) (domain.PaginatedSongTextResponse, error)
	Delete(songID int) error
	Update(songID int, input domain.SongUpdateInput) error
}

type Service struct {
	Song
}

func NewService(r *repository.Repository) *Service {
	return &Service{Song: NewSongService(r.Song)}
}
