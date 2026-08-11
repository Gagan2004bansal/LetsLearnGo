package sqlite

import (
	"database/sql"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO Students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	defer stmt.Close()
	res, errr := stmt.Exec(name, email, age)
	if errr != nil {
		return 0, errr
	}

	lastId, errror := res.LastInsertId()
	if errror != nil {
		return 0, err
	}

	return lastId, nil
}
