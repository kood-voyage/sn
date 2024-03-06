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

	if err := p.store.Privacy().Set(post.ID, privacy); err != nil {
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

func (p *PostRepository) GetUsers(source_id, target_id string) ([]model.Post, error) {
	query := `SELECT post.id, post.title, post.content, post.user_id, post.created_at
		FROM post
		JOIN user ON post.user_id = user.id
		JOIN privacy AS post_privacy ON post.id = post_privacy.id
		WHERE
    		user.id = ? AND (
			? = ? 
        	OR post_privacy.type_id = 1 -- Public
        	OR (
            	post_privacy.type_id = 2 -- Private
            		AND ? IN (
                	SELECT source_id
                	FROM follower
                	WHERE target_id = ?
           		)
        		)
        	OR (
            	post_privacy.type_id = 3 -- Selected
            	AND ? IN (
                	SELECT user_id
                	FROM selected_users
                	WHERE parent_id = post.id
            	)
        	)
    );`

	var posts []model.Post
	rows, err := p.store.Db.Query(query, target_id, source_id, target_id, source_id, target_id, source_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
