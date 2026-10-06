package players

import (
	"OCM/bot/lib/e"
	"OCM/bot/models"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
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
	query := `
			INSERT INTO players (
			first_name, 
			last_name, 
			patronymic, 
			phone_number,
			email,
			team_ids
			) VALUES (пше 
			$1, $2, $3, $4, $5, $6 
			)
			RETURNING id
	`

	var id int64

	err := s.db.QueryRow(
		query,
		p.FirstName,
		p.LastName,
		p.Patronymic,
		p.PhoneNumber,
		p.Email,
		p.TeamIDs,
	).Scan(&id)

	if err != nil {
		return 0, e.Wrap("can't save info", err)
	}

	return id, nil
}
