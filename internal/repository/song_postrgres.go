package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) Create(groupName string, songName string, s domain.Song) (int, error) {
	// tx, err := r.db.Begin()
	// if err != nil {
	// 	return 0, nil
	// }

	// queryString := fmt.Sprintf("INSERT INTO %s (user_id, title, msg, remind_date) VALUES ($1,$2,$3,$4) RETURNING id", remindTable)
	// row := tx.QueryRow(queryString, userID, rem.Title, rem.Msg, rem.RemindDate)
	// var id int
	// if err := row.Scan(&id); err != nil {
	// 	tx.Rollback()
	// 	return 0, err
	// }

	// return id, tx.Commit()
	return 0, nil
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
