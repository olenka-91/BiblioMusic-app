package service

import (
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
)

type SongService struct {
	repo repository.Song
}

func NewSongService(r repository.Song) *SongService {
	return &SongService{repo: r}
}

func (r *SongService) Create(s domain.SongList) (int, error) {
	return r.repo.Create(s)
}

func (r *SongService) GetSongsList(s domain.PaginatedSongInput) ([]domain.SongOutput, error) {
	return r.repo.GetSongsList(s)
}

func (r *SongService) GetSongText(s domain.PaginatedSongTextInput) (domain.PaginatedSongTextResponse, error) {
	return r.repo.GetSongText(s)
}

func (r *SongService) Delete(songID int) error {
	return r.repo.Delete(songID)
}

func (r *SongService) Update(songID int, input domain.SongUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return r.repo.Update(songID, input)
}
