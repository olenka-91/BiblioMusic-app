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

func (r *SongService) Create(req domain.AddSongRequest, detail domain.SongDetail) (int, error) {
	log.Debugf("Creating song...")
	s := domain.Song{
		GroupName:   &req.Group,
		Title:       &req.Song,
		Text:        detail.Text,
		ReleaseDate: detail.ReleaseDate,
		Link:        detail.Link}
	return r.repo.Create(s)
}

func (r *SongService) GetSongsList(s domain.PaginatedSongInput, page, pageSize string) ([]domain.Song, error) {
	log.Debugf("Fetching list of songs with pagination: %+v", s)
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
	log.Debugf("Fetching song text with pagination: %+v", s)
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
	log.Debugf("Deleting song with ID: %d", songID)
	return r.repo.Delete(songID)
}

func (r *SongService) Update(songID int, input domain.SongUpdateInput) error {
	log.Debugf("Updating song with ID: %d ...", songID)
	if err := input.Validate(); err != nil {
		log.Warn("Empty input data for song update")
		return err
	}
	return r.repo.Update(songID, input)
}
