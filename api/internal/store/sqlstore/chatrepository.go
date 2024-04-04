package sqlstore

import (
	"social-network/internal/model"

	"github.com/google/uuid"
)

type ChatRepository struct {
	store *Store
}

func (c *ChatRepository) Create(chat model.Chat) (*model.Chat, error) {
	query := `INSERT INTO chat id VALUES ?`

	_, err := c.store.Db.Exec(query, chat.ID)

	if err != nil {
		return nil, err
	}

	return &chat, err
}

func (c *ChatRepository) AddUser(user model.User, chat model.Chat) error {
	query := `INSERT INTO chat_users (id, user_id, chat_id) VALUES (?, ?, ?)`

	_, err := c.store.Db.Exec(query, uuid.New().String(), user.ID, chat.ID)
	if err != nil {
		return err
	}
	return err
}

func (c *ChatRepository) GetUsers(chat model.Chat) ([]model.User, error) {
	query := `SELECT user_id FROM chat_users WHERE chat_id = ?`

	row, err := c.store.Db.Query(query, chat.ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var users []model.User
	for row.Next() {
		var user model.User
		if err := row.Scan(&user.ID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (c *ChatRepository) AddLine(line *model.ChatLine) (*model.ChatLine, error) {
	query := `INSERT INTO chat_lines (id, chat_id, user_id, created_at, message) VALUES (?, ?, ?, ?, ?)
	`

	_, err := c.store.Db.Exec(query, line.ID, line.ChatID, line.UserID, line.CreatedAt, line.Message)
	if err != nil {
		return nil, err
	}
	return line, err
}
