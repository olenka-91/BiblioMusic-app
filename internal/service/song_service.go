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

func (r *SongService) Create(groupName string, songName string, s domain.Song) (int, error) {
	return r.repo.Create(groupName, songName, s)
}

// func (r *RemindService) GetByID(userID int, remindID int) (domain.Remind, error) {
// 	return r.repo.GetByID(userID, remindID)
// }
// func (r *RemindService) GetAll(userID int) ([]domain.Remind, error) {
// 	return r.repo.GetAll(userID)
// }
// func (r *RemindService) Delete(userID, remindID int) error {
// 	return r.repo.Delete(userID, remindID)
// }
// func (r *RemindService) Update(userID, remindID int, input domain.RemindUpdateInput) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}
// 	return r.repo.Update(userID, remindID, input)

// }
