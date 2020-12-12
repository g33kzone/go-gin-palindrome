package db

import (
	"fmt"
	"time"

	"github.com/g33kzone/go-message-palindrome/config"
	"github.com/g33kzone/go-message-palindrome/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //db utils
	log "github.com/sirupsen/logrus"
)

// Db - DB connection struct
type Db struct {
	dbConn *sqlx.DB
	err    error
}

//InitDBConnection - Initializes db connection
func (dbInfo *Db) InitDBConnection(pgConf config.PostgresConfig, attempt uint8) {
	log.Println(fmt.Sprintf("User %s connecting to db %s:%v", pgConf.DBUser, pgConf.DBServer, pgConf.DBPort))

	log.Printf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		pgConf.DBUser, pgConf.DBPwd, pgConf.DBServer, pgConf.DBPort, pgConf.DBName)

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			pgConf.DBUser, pgConf.DBPwd, pgConf.DBServer, pgConf.DBPort, pgConf.DBName))

	if err != nil {
		dbInfo.handleDBConnectionError(pgConf, attempt)
	}
	dbInfo.dbConn = db
	dbInfo.err = err
	log.Println(fmt.Sprintf("User %s connected to db %s:%v", pgConf.DBUser, pgConf.DBServer, pgConf.DBPort))
}

//handleDBConnectionError - Retires connecting to DB if attempts are not exhausted as per user defined limit
func (dbInfo *Db) handleDBConnectionError(pgConf config.PostgresConfig, attempt uint8) {
	if attempt < pgConf.DBRetryAttempts {
		time.Sleep(time.Millisecond * 500)
		dbInfo.InitDBConnection(pgConf, attempt+1)
	}
	log.Println("Connection Failed")
	panic(dbInfo.err)
}

//CloseDBConnection - Close DB connection if needed
func (dbInfo *Db) CloseDBConnection() {
	dbInfo.dbConn.Close()
}

//SaveMessage - Save messages in notes db
func (dbInfo *Db) SaveMessage(note model.Note) (uint, error) {
	var id uint = 0

	sqlStatement := `INSERT INTO post (message,ispalindrome) VALUES ($1,$2) RETURNING id`

	err := dbInfo.dbConn.QueryRow(sqlStatement, note.Message, note.IsPalindrome).Scan(&id)

	return id, err
}

//FetchAllMessages - Fetch all recent messages in notes db
func (dbInfo *Db) FetchAllMessages() ([]model.NoteResponse, error) {

	notesResponse := make([]model.NoteResponse, 0)

	rows, err := dbInfo.dbConn.Query("SELECT id, message,ispalindrome FROM post LIMIT $1", 50)
	if err != nil {
		return notesResponse, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint
		var message string
		var ispalindrome bool
		err := rows.Scan(&id, &message, &ispalindrome)
		if err != nil {
			return notesResponse, err
		}
		note := model.NoteResponse{ID: id, Message: message, IsPalindrome: ispalindrome}

		notesResponse = append(notesResponse, note)
	}
	return notesResponse, err
}

//FetchMessage - Fetch a single message in notes db
func (dbInfo *Db) FetchMessage(id uint) (model.NoteResponse, error) {
	var note model.NoteResponse
	rows, err := dbInfo.dbConn.Query("SELECT id, message,ispalindrome FROM post where id=$1", id)
	if err != nil {
		return note, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint
		var message string
		var ispalindrome bool
		err := rows.Scan(&id, &message, &ispalindrome)
		if err != nil {
			return note, err
		}
		note = model.NoteResponse{ID: id, Message: message, IsPalindrome: ispalindrome}
	}
	return note, err
}

//DeleteMessage - Delete a single message in notes db
func (dbInfo *Db) DeleteMessage(id uint) (int64, error) {

	sqlStatement := `DELETE FROM post WHERE id = $1;`

	sqlResult, err := dbInfo.dbConn.Exec(sqlStatement, id)
	count, _ := sqlResult.RowsAffected()

	return count, err
}
