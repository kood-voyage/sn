package sqlstore

import (
	"errors"
	"social-network/internal/model"

	"github.com/google/uuid"
)

type ChatRepository struct {
	store *Store
}

func (c *ChatRepository) Create(chat model.Chat) (*model.Chat, error) {
	query := `INSERT INTO chat (id, group_id) VALUES (?, ?)`

	_, err := c.store.Db.Exec(query, chat.ID, chat.GroupID)

	if err != nil {
		return nil, err
	}

	return &chat, err
}

func (c *ChatRepository) AddUser(user model.User, chat model.Chat) error {
	// Check if the entry already exists
	existsQuery := `SELECT COUNT(*) FROM chat_users WHERE user_id = ? AND chat_id = ?`
	var count int
	err := c.store.Db.QueryRow(existsQuery, user.ID, chat.ID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("entry already exists")
	}
	query := `INSERT INTO chat_users (id, user_id, chat_id) VALUES (?, ?, ?)`

	_, err = c.store.Db.Exec(query, uuid.New().String(), user.ID, chat.ID)
	if err != nil {
		return err
	}
	return err
}

func (c *ChatRepository) GetUsers(chat model.Chat) ([]model.User, error) {
	query := `SELECT u.id, u.username, u.email, u.password, u.timestamp, u.date_of_birth, u.first_name, u.last_name, u.description, u.avatar, u.cover
	FROM chat_users cu
	JOIN user u ON cu.user_id = u.id
	WHERE cu.chat_id = ?;
	`

	row, err := c.store.Db.Query(query, chat.ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var users []model.User
	for row.Next() {
		var user model.User
		if err := row.Scan(
            &user.ID,
            &user.Username,
            &user.Email,
            &user.Password,
            &user.CreatedAt,
            &user.DateOfBirth,
            &user.FirstName,
            &user.LastName,
            &user.Description,
            &user.Avatar,
            &user.Cover,
        ); err != nil {
            return nil, err
        }
		user.Sanitize()
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

func (c *ChatRepository) GetAll(user_id string) ([]*model.Chat, error) {
	query := `SELECT c.* FROM chat c JOIN chat_users cu ON c.id = cu.chat_id WHERE cu.user_id = ?`

	rows, err := c.store.Db.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*model.Chat
	for rows.Next() {
		var chat model.Chat
		if err := rows.Scan(&chat.ID, &chat.GroupID); err != nil {
			return nil, err
		}

		chats = append(chats, &chat)
	}

	return chats, nil
}

func (c *ChatRepository) Load(chat_id, user_id string) ([]*model.ChatLine, error) {
	query := `SELECT cl.* FROM chat_lines cl JOIN chat_users cu ON cl.chat_id = cu.chat_id WHERE cl.chat_id = ? AND cu.user_id = ? ORDER BY cl.created_at DESC`

	rows, err := c.store.Db.Query(query, chat_id, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatLines []*model.ChatLine
	for rows.Next() {
		var chatLine model.ChatLine
		if err := rows.Scan(&chatLine.ID, &chatLine.ChatID, &chatLine.UserID, &chatLine.CreatedAt, &chatLine.Message); err != nil {
			return nil, err
		}

		chatLines = append(chatLines, &chatLine)
	}

	return chatLines, nil
}

func (c *ChatRepository) GetChatsForUser(userID string) ([]*model.User, error) {
	query := `SELECT DISTINCT cu.user_id FROM chat_users cu JOIN chat_users cu2 ON cu.chat_id = cu2.chat_id WHERE cu2.user_id = ? AND cu.user_id != ?`

	rows, err := c.store.Db.Query(query, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID); err != nil {
			return nil, err
		}

		user.Sanitize()
		users = append(users, &user)
	}

	return users, nil
}
