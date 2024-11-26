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

func (r *SongPostgres) Create(s domain.Song) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	queryString := fmt.Sprintf("SELECT id FROM %s WHERE %s.name=$1", groupTable, groupTable)
	var GroupId int
	row := tx.QueryRow(queryString, s.GroupName)
	if err := row.Scan(&GroupId); err != nil {
		if err == sql.ErrNoRows {
			//если нет записи с такой группой-добавляем
			queryString = fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", groupTable)
			logrus.Debug("queryString=", queryString)
			logrus.Debug("s.Group=", s.GroupName)
			row = tx.QueryRow(queryString, s.GroupName)

			if err := row.Scan(&GroupId); err != nil {
				tx.Rollback()
				return 0, err
			}
		} else {
			tx.Rollback()
			return 0, err
		}
	}

	logrus.Debug("GroupId=", GroupId)
	queryString = fmt.Sprintf("INSERT INTO %s (group_id, title, text, release_date, link) VALUES ($1, $2, $3, $4, $5) RETURNING id", songTable)
	logrus.Debug("queryString=", queryString)
	logrus.Debug("s.Song=", domain.StringValue(s.Title))
	logrus.Debug("GroupId=", GroupId)
	logrus.Debug("s.Text=", domain.StringValue(s.Text))
	logrus.Debug("s.ReleaseDate=", domain.StringValue(s.ReleaseDate))
	logrus.Debug("s.Link=", domain.StringValue(s.Link))

	row = tx.QueryRow(queryString, GroupId, s.Title, s.Text, s.ReleaseDate, s.Link)
	var SongId int
	if err := row.Scan(&SongId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return SongId, tx.Commit()

}

func (r *SongPostgres) GetSongsList(s domain.PaginatedSongInput) ([]domain.Song, error) {
	offset := (s.Page - 1) * s.PageSize

	queryString := fmt.Sprintf(`SELECT %s.name as GroupName, title, release_date as ReleaseDate, text, link FROM %s 
								INNER JOIN %s ON %s.id=%s.group_id
								WHERE 1=1`,
		groupTable, songTable, groupTable, groupTable, songTable)

	args := make([]interface{}, 0)
	argCount := 1

	if s.GroupName != "" {
		queryString += fmt.Sprintf(" AND %s.name LIKE $%d", groupTable, argCount)
		args = append(args, "%"+s.GroupName+"%")
		argCount++
	}

	if s.Title != "" {
		queryString += fmt.Sprintf(" AND title LIKE $%d", argCount)
		args = append(args, "%"+s.Title+"%")
		argCount++
	}

	if s.ReleaseDate != "" {
		queryString += fmt.Sprintf(" AND release_date LIKE $%d", argCount)
		args = append(args, "%"+s.ReleaseDate+"%")
		argCount++
	}

	if s.Text != "" {
		queryString += fmt.Sprintf(" AND text LIKE $%d", argCount)
		args = append(args, "%"+s.Text+"%")
		argCount++
	}

	if s.Link != "" {
		queryString += fmt.Sprintf(" AND link LIKE $%d", argCount)
		args = append(args, "%"+s.Link+"%")
		argCount++
	}

	queryString += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, s.PageSize, offset)

	logrus.Debug("queryString=", queryString)
	logrus.Debug("args=", args)

	rows, err := r.db.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []domain.Song
	for rows.Next() {
		var s domain.Song
		if err := rows.Scan(&s.GroupName, &s.Title, &s.ReleaseDate, &s.Text, &s.Link); err != nil {
			logrus.Println("Error scanning row:", err)
			continue
		}
		songs = append(songs, s)
	}

	logrus.Debug("songs count=", len(songs))
	return songs, nil
}

func (r *SongPostgres) GetSongText(s domain.PaginatedSongTextInput) (domain.PaginatedSongTextResponse, error) {

	queryString := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", songTable)
	logrus.Debug("queryString=", queryString)
	logrus.Debug("s.SongId=", s.SongId)
	var song domain.SongDB
	err := r.db.Get(&song, queryString, s.SongId)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.PaginatedSongTextResponse{}, fmt.Errorf("song not found")
		}
		return domain.PaginatedSongTextResponse{}, err
	}

	verses := strings.Split(song.Text, "\\n\\n")
	logrus.Debug("len(verses)=", len(verses))

	offset := (s.Page - 1) * s.PageSize
	end := offset + s.PageSize
	if end > len(verses) {
		end = len(verses)
	}

	paginatedVerses := verses[offset:end]
	logrus.Debug("len(paginatedVerses)=", len(paginatedVerses))
	response := domain.PaginatedSongTextResponse{
		Title:       song.Title,
		Verses:      paginatedVerses,
		Page:        s.Page,
		PageSize:    s.PageSize,
		TotalVerses: len(verses),
	}

	return response, nil
}

func (r *SongPostgres) Delete(songID int) error {
	queryString := fmt.Sprintf("DELETE FROM %s t WHERE t.id=$1 RETURNING id", songTable)
	logrus.Debug("queryString=", queryString)
	logrus.Debug("songID=", songID)
	var deletedID int
	err := r.db.QueryRow(queryString, songID).Scan(&deletedID)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Warnf("No song with id=%d found to delete", songID)
		} else {
			logrus.Errorf("Failed to delete song with id=%d: %v", songID, err)
		}
	}

	return err
}

func (r *SongPostgres) Update(songID int, input domain.SongUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argIDs := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argIDs))
		args = append(args, *input.Title)
		argIDs++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argIDs))
		args = append(args, *input.Text)
		argIDs++
	}

	if input.ReleaseDate != nil {
		setValues = append(setValues, fmt.Sprintf("release_date=$%d", argIDs))
		args = append(args, *input.ReleaseDate)
		argIDs++
	}

	if input.Link != nil {
		setValues = append(setValues, fmt.Sprintf("link=$%d", argIDs))
		args = append(args, *input.Link)
		argIDs++
	}

	updateString := strings.Join(setValues, " ,")
	queryString := fmt.Sprintf("UPDATE %s t SET %s WHERE id = $%d", songTable, updateString, argIDs)
	args = append(args, songID)

	logrus.Debugf("updateQuery: %s", queryString)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(queryString, args...)
	if err != nil {
		logrus.Errorf("The song with id=%d not updated", songID)
	}

	return err
}
