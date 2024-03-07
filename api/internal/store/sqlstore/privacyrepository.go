package sqlstore

import (
	"database/sql"
	"errors"
)

type PrivacyRepository struct {
	store *Store
}

func (p *PrivacyRepository) Set(parent_id string, privacy int) error {
	query := `INSERT INTO privacy (id, type_id) VALUES (?,?)`

	_, err := p.store.Db.Exec(query, parent_id, privacy)
	if err != nil {
		return err
	}

	return nil
}

func (u *PrivacyRepository) Update(parent_id string, privacy int) error {
	query := `UPDATE privacy SET type_id = ? WHERE id = ?`

	result, err := u.store.Db.Exec(query, privacy, parent_id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (u *PrivacyRepository) Delete(parent_id string) error {
	query := `DELETE FROM privacy WHERE id = ?`

	result, err := u.store.Db.Exec(query, parent_id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (u *PrivacyRepository) Check(parent_id string) (int, error) {
	query := `SELECT type_id FROM privacy WHERE id = ?`

	var privacy int
	if err := u.store.Db.QueryRow(query, parent_id).Scan(&privacy); err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("privacy entry does not exist")
		}
		return -1, err
	}

	return privacy, nil
}
