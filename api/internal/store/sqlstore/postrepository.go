package sqlstore

import (
	"social-network/internal/model"
)

type PostRepository struct {
	store *Store
}

func (p PostRepository) Create(post *model.Post) error {
	query := `INSERT INTO post (
                  id,
                  user_id,
                  title,
                  content,
                  created_at) VALUES (?,?,?,?,?)`

	_, err := p.store.Db.Exec(query,
		post.ID,
		post.UserID,
		post.Title,
		post.Content,
		post.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p PostRepository) Delete(id string) error {
	query := `DELETE FROM post WHERE id = ?`

	_, err := p.store.Db.Exec(query, id)
	if err != nil {
		return err
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
