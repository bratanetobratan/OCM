package teams

import (
	"OCM/bot/lib/e"
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

func (s *Storage) Join(teamID int64, playerID int64) error {
	query := `
            INSERT INTO team_rosters (
                team_id,
                player_id
            ) VALUES (
                $1, $2 
            )
            ON CONFLICT (team_id, player_id) DO NOTHING
    `

	_, err := s.db.Exec(
		query,
		teamID,
		playerID,
	)

	if err != nil {
		return e.Wrap("can't add player to team_rosters", err)
	}

	return nil
}

func (s *Storage) Request(teamID int64, playerID int64) (ID int64, err error) {
	defer func() { err = e.WrapIfErr("can't create team request", err) }()

	query := `
            INSERT INTO team_requests (
                team_id,
                player_id,
                status
            ) VALUES (
                $1, $2, 'pending'
            )
            ON CONFLICT (team_id, player_id, status) DO NOTHING
            RETURNING id
    `

	var requestID int64

	err = s.db.QueryRow(
		query,
		teamID,
		playerID,
	).Scan(&requestID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}

	return requestID, nil
}

func (s *Storage) Reply(requestID int64, approved bool) (err error) {
	defer func() { err = e.WrapIfErr("can't process coach reply", err) }()

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	status := "rejected"
	if approved {
		status = "approved"
	}

	var teamID, playerID int64

	query := `
            UPDATE team_requests
            SET status = $1
            WHERE id = $2 AND status = 'pending'
            RETURNING team_id, player_id
    `

	err = tx.QueryRow(
		query,
		status,
		requestID,
	).Scan(&teamID, &playerID)

	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}

	if approved {
		query = `
                INSERT INTO team_rosters (
                    team_id,
                    player_id
                ) VALUES (
                    $1, $2 
                )
                ON CONFLICT (team_id, player_id) DO NOTHING
        `

		_, err = tx.Exec(
			query,
			teamID,
			playerID,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
