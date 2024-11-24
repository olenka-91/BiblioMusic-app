package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/sirupsen/logrus"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) Create(s domain.SongList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	queryString := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", groupTable)
	row := tx.QueryRow(queryString, s.Group)
	var GroupId int
	if err := row.Scan(&GroupId); err != nil {
		tx.Rollback()
		return 0, err
	}

	queryString = fmt.Sprintf("INSERT INTO %s (group_id, title) VALUES ($1,$2) RETURNING id", songTable)
	row = tx.QueryRow(queryString, GroupId, s.Song)
	var SongId int
	if err := row.Scan(&SongId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return SongId, tx.Commit()
}

func (r *SongPostgres) GetSongsList(s domain.PaginatedSongInput) ([]domain.SongOutput, error) {
	offset := (s.Page - 1) * s.PageSize

	queryString := fmt.Sprintf(`SELECT %s.name as GroupName, title, release_date as ReleaseDate, text, link FROM %s 
								INNER JOIN %s ON %s.id=%s.group_id
								WHERE %s.name LIKE $1 
								AND title LIKE $2
								AND release_date LIKE $3
								AND text LIKE $4
								AND link LIKE $5
								LIMIT $6 OFFSET $7`,
		groupTable, songTable, groupTable, groupTable, songTable, groupTable)

	logrus.Debug("queryString=", queryString, "%"+s.GroupName+"%", "%"+s.Title+"%", s.PageSize, "offset=", offset)

	rows, err := r.db.Query(queryString,
		"%"+s.GroupName+"%",
		"%"+s.Title+"%",
		"%"+s.ReleaseDate+"%",
		"%"+s.Text+"%",
		"%"+s.Link+"%",
		s.PageSize,
		offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []domain.SongOutput
	for rows.Next() {
		var s domain.SongOutput
		if err := rows.Scan(&s.GroupName, &s.Title, &s.ReleaseDate, &s.Text, &s.Link); err != nil {
			logrus.Println("Error scanning row:", err)
			continue
		}
		songs = append(songs, s)
	}

	return songs, nil
}

func (r *SongPostgres) GetSongText(s domain.PaginatedSongTextInput) (domain.PaginatedSongTextResponse, error) {

	queryString := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", songTable)

	var song domain.Song
	err := r.db.Get(&song, queryString, s.SongId)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.PaginatedSongTextResponse{}, fmt.Errorf("song not found")
		}
		return domain.PaginatedSongTextResponse{}, err
	}

	verses := strings.Split(song.Text, "\\n\\n")

	offset := (s.Page - 1) * s.PageSize
	end := offset + s.PageSize
	if end > len(verses) {
		end = len(verses)
	}

	paginatedVerses := verses[offset:end]

	response := domain.PaginatedSongTextResponse{
		Title:       song.Title,
		Verses:      paginatedVerses,
		Page:        s.Page,
		PageSize:    s.PageSize,
		TotalVerses: len(verses),
	}

	return response, nil
}

// func (r *RemindPostgres) GetByID(userID int, remindID int) (domain.Remind, error) {
// 	var rem domain.Remind
// 	queryString := fmt.Sprintf("SELECT t.id, t.title, t.msg, t.remind_date as RemindDate FROM %s t WHERE t.id=$1 and t.user_id=$2", remindTable)
// 	err := r.db.Get(&rem, queryString, remindID, userID)

// 	return rem, err
// }

// func (r *RemindPostgres) GetAll(userID int) ([]domain.Remind, error) {
// 	var rem []domain.Remind
// 	queryString := fmt.Sprintf("SELECT t.id, t.title, t.msg, t.remind_date as RemindDate FROM %s t WHERE t.user_id=$1", remindTable)
// 	//logrus.Debug("queryString=", queryString, " userID=", userID)
// 	err := r.db.Select(&rem, queryString, userID)

// 	return rem, err
// }

// func (r *RemindPostgres) Delete(userID, remindID int) error {
// 	queryString := fmt.Sprintf("DELETE FROM %s t WHERE t.id=$1 and t.user_id=$2", remindTable)
// 	_, err := r.db.Exec(queryString, remindID, userID)
// 	return err
// }

// func (r *RemindPostgres) Update(userID, remindID int, input domain.RemindUpdateInput) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argIDs := 1

// 	if input.Title != nil {
// 		setValues = append(setValues, fmt.Sprintf("title=$%d", argIDs))
// 		args = append(args, *input.Title)
// 		argIDs++
// 	}

// 	if input.Msg != nil {
// 		setValues = append(setValues, fmt.Sprintf("msg=$%d", argIDs))
// 		args = append(args, *input.Msg)
// 		argIDs++
// 	}

// 	if input.RemindDate != nil {
// 		setValues = append(setValues, fmt.Sprintf("remind_date=$%d", argIDs))
// 		args = append(args, *input.RemindDate)
// 		argIDs++
// 	}

// 	updateString := strings.Join(setValues, " ,")
// 	queryString := fmt.Sprintf("UPDATE %s t SET %s WHERE id = $%d AND user_id=$%d", remindTable, updateString, argIDs, argIDs+1)
// 	args = append(args, remindID, userID)

// 	logrus.Debugf("updateQuery: %s", queryString)
// 	logrus.Debugf("args: %s", args)

// 	_, err := r.db.Exec(queryString, args...)

// 	return err
// }
