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
			) VALUES ( 
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

func (s *Storage) IsPhoneExist(phoneNumber string) (bool, error) {
	query := `SELECT EXISTS(
				SELECT 1
				FROM players
				WHERE phone_number = $1			
			)`

	var exists bool

	err := s.db.QueryRow(
		query,
		phoneNumber,
	).Scan(&exists)

	if err != nil {
		return false, e.Wrap("failed to check phone existence", err)
	}

	return exists, nil
}

func (s *Storage) IsEmailExist(email string) (bool, error) {
	query := `SELECT EXISTS(
				SELECT 1
				FROM players
				WHERE email = $1
			)`

	var exists bool

	err := s.db.QueryRow(
		query,
		email,
	).Scan(&exists)

	if err != nil {
		return false, e.Wrap("failed to check email existence", err)
	}

	return exists, nil
}
