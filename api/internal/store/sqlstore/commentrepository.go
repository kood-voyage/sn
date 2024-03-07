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

func (c CommentRepository) Delete(commentID, userID string) error {
	query := `DELETE FROM comment WHERE id = ? AND user_id = ?`

	result, err := c.store.Db.Exec(query, commentID, userID)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no comment with such ID %s", commentID)
	}

	return nil
}

// Get returns all comments to single post
func (c CommentRepository) Get(id string) (*[]model.Comment, error) {
	//query := `SELECT * FROM comment WHERE post_id = ?`
	q := `WITH RECURSIVE CommentHierarchy AS (
        -- Anchor member: Start with the top-level comments for the post
        SELECT
            c.id,
            c.user_id,
            c.post_id,
            c.parent_id,
            c.content,
            c.timestamp,
            (SELECT COUNT(*) FROM comment subc WHERE subc.parent_id = c.id) AS subcomment_count
        FROM
            comment c
        WHERE
            c.post_id = ? AND (c.parent_id IS NULL OR c.parent_id = '')
    
        UNION ALL

        -- Recursive member: Join with sub-comments
        SELECT
            c.id,
            c.user_id,
            c.post_id,
            c.parent_id,
            c.content,
            c.timestamp,
            (SELECT COUNT(*) FROM comment subc WHERE subc.parent_id = c.id) AS subcomment_count
        FROM
            comment c
        JOIN
            CommentHierarchy ch ON c.parent_id = ch.id
    )
    SELECT * FROM CommentHierarchy;`
	rows, err := c.store.Db.Query(q, id)
	if err != nil {
		return nil, err
	}

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err = rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.ParentID,
			&comment.Content,
			&comment.CreatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &comments, nil
}
