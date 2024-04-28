package sqlstore

import (
	"errors"
	"social-network/internal/model"
)

type SessionRepository struct {
	store *Store
}

func (s *SessionRepository) Create(session model.Session) (*model.Session, error) {
	query := `INSERT INTO session (access_id, user_id, timestamp) VALUES (?, ?, ?)`

	_, err := s.store.Db.Exec(query, session.AcessID, session.UserID, session.CreatedAT)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionRepository) CheckByUserId(user_id string) (*model.Session, error) {
	query := `SELECT * FROM session WHERE user_id = ?`

	var session model.Session

	if err := s.store.Db.QueryRow(query, user_id).Scan(&session.AcessID, &session.UserID, &session.CreatedAT); err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionRepository) Update(access_id string, session model.Session) (*model.Session, error) {
	query := `UPDATE session SET access_id = ?, user_id = ?, timestamp = ? WHERE access_id = ?`

	rows, err := s.store.Db.Exec(query, session.AcessID, session.UserID, session.CreatedAT, access_id)
	if err != nil {
		return nil, err
	}

	aff, err := rows.RowsAffected()
	if err != nil {
		return nil, err
	}
	if aff == 0 {
		return nil, errors.New("no rows affected")
	}

	return &session, nil
}

func (s *SessionRepository) Delete(access_id string) error {
	query := `DELETE FROM session WHERE access_id = ?`

	rows, err := s.store.Db.Exec(query, access_id)
	if err != nil {
		return err
	}

	aff, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (s *SessionRepository) Check(access_id string) (*model.Session, error) {
	query := `SELECT * FROM session WHERE access_id = ?`

	var session model.Session
	if err := s.store.Db.QueryRow(query, access_id).Scan(&session.AcessID, &session.UserID, &session.CreatedAT); err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionRepository) DeleteByUser(user_id string) error {
	query := `DELETE FROM session WHERE user_id = ?`

	rows, err := s.store.Db.Exec(query, user_id)
	if err != nil {
		return err
	}

	aff, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("no rows affected")
	}

	return nil
}