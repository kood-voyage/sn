package sqlstore

import (
	"social-network/internal/model"
)

type UserRepository struct {
	store *Store
}

func (u *UserRepository) Create(user *model.User, privacy int) error {
	query := `INSERT INTO user (id) VALUES (?)`

	_, err := u.store.Db.Exec(query, user.ID)

	if err != nil {
		return err
	}

	//insert user privacy state to database
	query = `INSERT INTO privacy (id, type_id) VALUES (?, ?)`

	_, err = u.store.Db.Exec(query, user.ID, privacy)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdatePrivacy(user *model.User, privacy int) error {
	query := `UPDATE privacy SET type_id = ? WHERE id = ?`

	_, err := u.store.Db.Exec(query, privacy, user.ID)

	if err != nil {
		return err
	}

	return nil
}
