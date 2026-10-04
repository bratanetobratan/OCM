package players

import (
	"OCM/bot/models"
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Save(p *models.Player) (int64, error) {
	
}
