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
                     content,
                     user_name,
                     user_avatar) VALUES (?,?,?,?,?,?,?)`

	_, err := c.store.Db.Exec(query,
		comment.ID,
		comment.UserID,
		comment.PostID,
		getParentID(comment),
		comment.Content,
		comment.UserName,
		comment.UserAvatar)
	if err != nil {
		return err
	}

	if err = c.store.Image().Add(comment.ID, comment.ImagePaths); err != nil {
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

	if err = c.store.Image().DeleteAll(commentID); err != nil {
		return err
	}

	return nil
}
func (c CommentRepository) GetAll(id string) (*[]model.Comment, error) {
	q := `
    WITH RECURSIVE CommentHierarchy AS (
		SELECT
			c.id,
			c.user_id,
			c.post_id,
			COALESCE(c.parent_id, '') AS parent_id,
			c.content,
			c.created_at,
			c.user_name,
			c.user_avatar,
			(SELECT COUNT(*) FROM comment subc WHERE subc.parent_id = c.id) AS count
		FROM
			comment c
		WHERE
			c.post_id = ? AND (c.parent_id IS NULL OR c.parent_id = '')
		
		UNION ALL
	
		SELECT
			c.id,
			c.user_id,
			c.post_id,
			COALESCE(c.parent_id, '') AS parent_id,
			c.content,
			c.created_at,
			c.user_name,
			c.user_avatar,
			(SELECT COUNT(*) FROM comment subc WHERE subc.parent_id = c.id) AS count
		FROM
			comment c
		JOIN
			CommentHierarchy ch ON c.parent_id = ch.id
	)
	SELECT
		ch.*,
		u.*
	FROM
		CommentHierarchy ch
	LEFT JOIN
		user u ON ch.user_id = u.id;
    `

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
			&comment.UserName,
			&comment.UserAvatar,
			&comment.Count,
			&comment.UserInformation.ID,
			&comment.UserInformation.Username,
			&comment.UserInformation.Email,
			&comment.UserInformation.Password,
			&comment.UserInformation.CreatedAt,
			&comment.UserInformation.DateOfBirth,
			&comment.UserInformation.FirstName,
			&comment.UserInformation.LastName,
			&comment.UserInformation.Description,
			&comment.UserInformation.Avatar,
			&comment.UserInformation.Cover,
		); err != nil {
			return nil, err
		}
		comment.UserInformation.Sanitize()
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &comments, nil
}

func (c CommentRepository) Update(comment *model.Comment) error {
	query := `UPDATE comment SET content = ? WHERE id = ?`

	_, err := c.store.Db.Exec(query, comment.Content, comment.ID)
	if err != nil {
		return err
	}

	if err = c.store.Image().Update(comment.ID, comment.ImagePaths); err != nil {
		return err
	}
	return nil
}

func (c CommentRepository) IsAuthor(comment *model.Comment, userId string) bool {
	query := `SELECT user_id FROM comment WHERE id = ?`

	var storedUserID string
	if err := c.store.Db.QueryRow(query, comment.ID).Scan(&storedUserID); err != nil {
		return false
	}

	return storedUserID == userId
}
