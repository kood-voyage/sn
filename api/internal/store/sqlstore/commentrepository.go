package sqlstore

import (
	"fmt"
	"social-network/internal/model"
)

type CommentRepository struct {
	store *Store
}

func (c CommentRepository) Create(comment *model.Comment) error {
	query := `INSERT INTO comment (
                     id,
                     user_id,
                     post_id,
                     parent_id,
                     content) VALUES (?,?,?,?,?)`

	_, err := c.store.Db.Exec(query,
		comment.ID,
		comment.UserID,
		comment.PostID,
		getParentID(comment),
		comment.Content)
	if err != nil {
		return err
	}

	return nil
}

func getParentID(comment *model.Comment) interface{} {
	if comment.ParentID == "" {
		return nil
	}
	return comment.ParentID
}

func (c CommentRepository) Delete(id string) error {
	query := `DELETE FROM comment WHERE id = ?`

	result, err := c.store.Db.Exec(query, id)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no comment with such ID %s", id)
	}

	return nil
}

// Get returns all comments to single post
func (c CommentRepository) Get(id string) (*model.Comment, error) {

	return nil, nil
}
