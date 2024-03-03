package sqlstore

import "social-network/internal/model"

type CommentRepository struct {
	store *Store
}

func (c CommentRepository) Create(comment *model.Comment) error {

	return nil
}

func (c CommentRepository) Delete(id string) error {
	query := `DELETE FROM post WHERE id = ?`

	_, err := c.store.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (c CommentRepository) Get(id string) (*model.Comment, error) {

	return nil, nil
}
