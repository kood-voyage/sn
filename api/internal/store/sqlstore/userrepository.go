package sqlstore

import (
	"database/sql"
	"errors"
	"social-network/internal/model"
)

type UserRepository struct {
	store *Store
}

func (u *UserRepository) Create(user model.User, privacy int) (*model.User, error) {
	query := `INSERT INTO user (id, username, email, password, date_of_birth, first_name, last_name, avatar, cover, description) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?)`

	if err := user.BeforeCreate(); err != nil {
		return nil, err
	}

	_, err := u.store.Db.Exec(query, user.ID, user.Username, user.Email, user.Password, user.DateOfBirth, user.FirstName, user.LastName, user.Avatar, user.Cover, user.Description)
	if err != nil {
		return nil, err
	}

	//insert user privacy state to database
	if err := u.store.Privacy().Set(user.ID, privacy); err != nil {
		return nil, err
	}

	user.Sanitize()
	return &user, nil
}

func (u *UserRepository) GetFollowers(userID string) ([]model.User, error) {
	query := `SELECT u.* 
	FROM follower f 
	JOIN user u ON f.source_id = u.id 
	WHERE f.target_id = ?`

	var followers []model.User

	rows, err := u.store.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.DateOfBirth, &user.FirstName, &user.LastName, &user.Description, &user.Avatar, &user.Cover); err != nil {
			return nil, err
		}
		user.Sanitize()
		followers = append(followers, user)
	}

	return followers, nil
}

func (u *UserRepository) GetFollowing(userID string) ([]model.User, error) {
	query := `SELECT u.*
	FROM follower f
	JOIN user u ON f.target_id = u.id
	WHERE f.source_id = ?`

	var followers []model.User

	rows, err := u.store.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.DateOfBirth, &user.FirstName, &user.LastName, &user.Description, &user.Avatar, &user.Cover); err != nil {
			return nil, err
		}
		user.Sanitize()
		followers = append(followers, user)
	}

	return followers, nil
}

func (u *UserRepository) IsFollowing(source_id, target_id string) (bool, error) {
	query := `SELECT id FROM follower WHERE source_id = ? AND target_id = ?`
	var target string
	if err := u.store.Db.QueryRow(query, source_id, target_id).Scan(&target); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *UserRepository) GetNotifications(user_id string) ([]model.Request, error) {
	query := `SELECT 
    request.*, 
    source_user.*, 
    target_user.*
FROM 
    request
JOIN 
    user AS source_user ON request.source_id = source_user.id
JOIN 
    user AS target_user ON request.target_id = target_user.id
WHERE 
    target_id = ?;
`

	var notifications []model.Request

	rows, err := u.store.Db.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notification model.Request
		var sourceUser model.User
		var targetUser model.User
		if err := rows.Scan(
			&notification.ID,
			&notification.TypeID,
			&notification.SourceID,
			&notification.TargetID,
			&notification.ParentID,
			&notification.Message,
			&notification.CreatedAt,
			&sourceUser.ID,
			&sourceUser.Username,
			&sourceUser.Email,
			&sourceUser.Password,
			&sourceUser.CreatedAt,
			&sourceUser.DateOfBirth,
			&sourceUser.FirstName,
			&sourceUser.LastName,
			&sourceUser.Description,
			&sourceUser.Avatar,
			&sourceUser.Cover,
			&targetUser.ID,
			&targetUser.Username,
			&targetUser.Email,
			&targetUser.Password,
			&targetUser.CreatedAt,
			&targetUser.DateOfBirth,
			&targetUser.FirstName,
			&targetUser.LastName,
			&targetUser.Description,
			&targetUser.Avatar,
			&targetUser.Cover,
		); err != nil {
			return nil, err
		}
		notification.SourceInformation = sourceUser
		notification.TargetInformation = targetUser
		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (u *UserRepository) GetAll() ([]model.User, error) {
	query := `SELECT * FROM user`

	var users []model.User

	rows, err := u.store.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.DateOfBirth, &user.FirstName, &user.LastName, &user.Description, &user.Avatar, &user.Cover); err != nil {
			return nil, err
		}
		user.Sanitize()
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) SetDescription(user_id, description string) error {
	query := `UPDATE user SET description = ? WHERE id = ?`
	row, err := u.store.Db.Exec(query, description, user_id)
	if err != nil {
		return err
	}

	aff, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (u *UserRepository) SetCover(user_id, cover string) error {
	query := `UPDATE user SET cover = ? WHERE id = ?`
	row, err := u.store.Db.Exec(query, cover, user_id)
	if err != nil {
		return err
	}

	aff, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (u *UserRepository) SetAvatar(user_id, avatar string) error {
	query := `UPDATE user SET avatar = ? WHERE id = ?`
	row, err := u.store.Db.Exec(query, avatar, user_id)
	if err != nil {
		return err
	}

	aff, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (u *UserRepository) Get(user_id string) (*model.User, error) {
	query := `SELECT * FROM user WHERE username = ? OR email = ?`

	var user model.User
	if err := u.store.Db.QueryRow(query, user_id, user_id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.DateOfBirth, &user.FirstName, &user.LastName, &user.Description, &user.Avatar, &user.Cover); err != nil {
		return nil, err
	}

	return &user, nil
}

