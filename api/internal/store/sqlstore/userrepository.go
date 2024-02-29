package sqlstore

import (
	"github.com/google/uuid"
	"social-network/internal/model"
	"time"
)

type UserRepository struct {
	store *Store
}

func (u *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO user (
                  id,
                  username,
                  email,
                  password,
                  created_at,
                  date_of_birth,
                  first_name,
                  last_name,
                  gender,
                  description) VALUES  (?,?,?,?,?,?,?,?,?,?)`

	user, err := prepareUser(user)
	if err != nil {
		return err
	}

	_, err = u.store.Db.Exec(
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.DateOfBirth,
		user.FirstName,
		user.LastName,
		user.Gender,
		user.Description)
	if err != nil {
		return err
	}

	user.Sanitize()

	return nil
}

func prepareUser(user *models.User) (*models.User, error) {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	err := user.BeforeCreate()
	if err != nil {
		return nil, err
	}

	return user, nil
}
