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

func (s *Storage) Save(p *models.Player, platform models.Platform, externalID string) (ID int64, err error) {
	defer func() { err = e.WrapIfErr("can't save player", err) }()

	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `
			INSERT INTO players (
				first_name, 
				last_name, 
				patronymic, 
				phone_number,
				email
			) VALUES ( 
				$1, $2, $3, $4, $5
			)
			RETURNING id
	`

	err = tx.QueryRow(
		query,
		p.FirstName,
		p.LastName,
		p.Patronymic,
		p.PhoneNumber,
		p.Email,
	).Scan(&ID)

	if err != nil {
		return 0, err
	}

	query = `
			INSERT INTO player_messengers (
				player_id,
				platform,
				external_id
			) VALUES (
				$1, $2, $3 
			)
	`

	_, err = tx.Exec(
		query,
		ID,
		string(platform),
		externalID,
	)

	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (s *Storage) Remove(id int64) error {
	_, err := s.db.Exec(
		`DELETE FROM players WHERE id = $1`,
		id,
	)

	if err != nil {
		return e.Wrap("can't remove player", err)
	}

	return nil
}

func (s *Storage) Update(p *models.Player) error {
	query := `
			UPDATE players
				SET first_name = $1,
				last_name = $2,
				patronymic = $3,
				phone_number = $4,
				email = $5
			WHERE id = $6
	`

	_, err := s.db.Exec(
		query,
		p.FirstName,
		p.LastName,
		p.Patronymic,
		p.PhoneNumber,
		p.Email,
		p.ID,
	)

	if err != nil {
		return e.Wrap("can't update player info", err)
	}

	return nil
}

func (s *Storage) Find(platform models.Platform, externalID string) (*models.Player, error) {
	query := `
			SELECT 
            	p.id, 
            	p.first_name, 
            	p.last_name, 
            	COALESCE(p.patronymic, ''), 
            	p.phone_number, 
            	p.email
        	FROM players p
        	JOIN player_messengers pm ON p.id = pm.player_id
        	WHERE pm.platform = $1 AND pm.external_id = $2
	`

	var p models.Player

	err := s.db.QueryRow(
		query,
		string(platform),
		externalID,
	).Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Patronymic,
		&p.PhoneNumber,
		&p.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, e.Wrap("can't find player", err)
	}

	return &p, nil
}

func (s *Storage) IsPhoneExist(phoneNumber string) (bool, error) {
	query := `
			SELECT EXISTS(
				SELECT 1
				FROM players
				WHERE phone_number = $1			
			)
	`

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
	query := `
			SELECT EXISTS(
				SELECT 1
				FROM players
				WHERE email = $1
			)
	`

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
