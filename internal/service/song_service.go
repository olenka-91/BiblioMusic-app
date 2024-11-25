package service

import (
	"strconv"

	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
	log "github.com/sirupsen/logrus"
)

const (
	defPage     = 1
	defPageSize = 10
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

func (r *SongService) GetSongsList(s domain.PaginatedSongInput, page, pageSize string) ([]domain.SongOutput, error) {
	log.Debug("Fetching list of songs with pagination")
	iPage, err := strconv.Atoi(page)
	if err != nil || iPage < 1 {
		log.Warn("Invalid page number, defaulting to page ", defPage)
		iPage = defPage
	}

	iPageSize, err := strconv.Atoi(pageSize)
	if err != nil || iPageSize < 1 {
		log.Warn("Invalid page size, defaulting to page size ", defPageSize)
		iPageSize = defPageSize
	}

	s.Page = iPage
	s.PageSize = iPageSize

	return r.repo.GetSongsList(s)
}

func (r *SongService) GetSongText(s domain.PaginatedSongTextInput, page, pageSize string) (domain.PaginatedSongTextResponse, error) {
	log.Debug("Fetching song text with pagination")
	iPage, err := strconv.Atoi(page)
	if err != nil || iPage < 1 {
		log.Warn("Invalid page number, defaulting to page ", defPage)
		iPage = defPage
	}

	iPageSize, err := strconv.Atoi(pageSize)
	if err != nil || iPageSize < 1 {
		log.Warn("Invalid page size, defaulting to page size ", defPageSize)
		iPageSize = defPageSize
	}

	s.Page = iPage
	s.PageSize = iPageSize

	return r.repo.GetSongText(s)
}

func (r *SongService) Delete(songID int) error {
	return r.repo.Delete(songID)
}

func (r *SongService) Update(songID int, input domain.SongUpdateInput) error {
	if err := input.Validate(); err != nil {
		log.Warn("Empty input data for song update")
		return err
	}
	return r.repo.Update(songID, input)
}
