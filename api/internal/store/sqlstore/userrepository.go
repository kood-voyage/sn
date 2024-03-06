package sqlstore

import (
	"database/sql"
	"errors"
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

func (u *UserRepository) CheckPrivacy(userID string) (int, error) {
	query := `SELECT type_id FROM privacy WHERE id = ?`

	var privacy int
	if err := u.store.Db.QueryRow(query, userID).Scan(&privacy); err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("user does not exist")
		}
		return -1, err
	}

	return privacy, nil
}

func (u *UserRepository) GetFollowers(userID string) ([]model.User, error) {
	query := `SELECT source_id FROM follower WHERE target_id = ?`

	var followers []model.User

	rows, err := u.store.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var follower model.User
		if err = rows.Scan(&follower.ID); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (u *UserRepository) GetFollowing(userID string) ([]model.User, error) {
	query := `SELECT target_id FROM follower WHERE source_id = ?`

	var followers []model.User

	rows, err := u.store.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var follower model.User
		if err = rows.Scan(&follower.ID); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (u *UserRepository) IsFollowing(source_id, target_id string) (bool, error) {
	query := `SELECT id FROM follower WHERE source_id = ? AND target_id = ?`
	var target string
	if err := u.store.Db.QueryRow(query, source_id, target_id).Scan(&target);err != nil{
		if err == sql.ErrNoRows{
			return false, nil
		} 
		return false, err
	}
	return true, nil
}

func (u *UserRepository) GetNotifications(user_id string, req_type int) ([]model.Request, error) {
	query := `SELECT * FROM request WHERE target_id = ? AND type_id = ?`

	var notifications []model.Request

	rows, err := u.store.Db.Query(query, user_id, req_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var notification model.Request
		if err := rows.Scan(&notification.ID, &notification.TypeID, &notification.SourceID, &notification.TargetID, &notification.Message, &notification.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
