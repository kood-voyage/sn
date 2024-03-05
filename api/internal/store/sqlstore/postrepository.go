package sqlstore

import (
	"fmt"
	"social-network/internal/model"
)

type PostRepository struct {
	store *Store
}

func (p PostRepository) Create(post *model.Post, privacy int) error {
	query := `INSERT INTO post (
                  id,
                  user_id,
                  title,
                  content) VALUES (?,?,?,?)`

	_, err := p.store.Db.Exec(query,
		post.ID,
		post.UserID,
		post.Title,
		post.Content)
	if err != nil {
		return err
	}
	query = `INSERT INTO privacy (id, type_id) VALUES (?,?)`

	_, err = p.store.Db.Exec(query, post.ID, privacy)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepository) Delete(id string) error {
	query := `DELETE FROM post WHERE id = ?`

	result, err := p.store.Db.Exec(query, id)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no post with such ID %s", id)
	}

	return nil
}

func (p PostRepository) Get(id string) (*model.Post, error) {
	post := &model.Post{}
	query := `SELECT * FROM post WHERE id = ?`

	err := p.store.Db.QueryRow(query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return post, nil
}
