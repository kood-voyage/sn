package sqlstore

import (
	"database/sql"
	"fmt"
	"social-network/internal/model"
	"strings"

	"github.com/google/uuid"
)

type PostRepository struct {
	store *Store
}

func (p *PostRepository) Create(post *model.Post, privacy int) error {
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

	if err = p.store.Privacy().Set(post.ID, privacy); err != nil {
		return err
	}

	if err = p.store.Image().Add(post.ID, post.ImagePaths); err != nil {
		return err
	}

	return nil
}

func (p *PostRepository) Update(post *model.Post, privacy int) error {
	query := `UPDATE post SET title = ?, content = ? WHERE id = ?`

	_, err := p.store.Db.Exec(query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	if err = p.store.Privacy().Update(post.ID, privacy); err != nil {
		return err
	}

	if err = p.store.Image().Update(post.ID, post.ImagePaths); err != nil {
		return err
	}

	return nil
}

func (p *PostRepository) Delete(id string) error {
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

	if err = p.store.Image().DeleteAll(id); err != nil {
		return err
	}

	return nil
}

func (p *PostRepository) Get(id string) (*model.Post, error) {
	post := &model.Post{}
	query := `SELECT 
                 p.id,
                 p.title,
                 p.content,
                 p.user_id,
                 COALESCE(p.community_id, '') AS community_id,
                 p.created_at,
                 i.path
             FROM 
                 post p
             LEFT JOIN 
                 image i ON p.id = i.parent_id
             WHERE 
                 p.id = ?`

	rows, err := p.store.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []string

	for rows.Next() {
		var path sql.NullString
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserID,
			&post.CommunityID,
			&post.CreatedAt,
			&path,
		)
		if err != nil {
			return nil, err
		}

		if path.Valid {
			paths = append(paths, path.String)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	post.ImagePaths = paths
	return post, nil
}

func (p *PostRepository) GetUsers(source_id, target_id string) ([]model.Post, error) {
	query := `SELECT 
    post.id,
    post.title,
    post.content,
    post.user_id,
    post.created_at,
    image.path
FROM 
    post
JOIN 
    user ON post.user_id = user.id
JOIN 
    privacy AS post_privacy ON post.id = post_privacy.id
LEFT JOIN 
    image ON post.id = image.parent_id
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
    )
GROUP BY
    post.id`

	var posts []model.Post
	rows, err := p.store.Db.Query(query, target_id, source_id, target_id, source_id, target_id, source_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var path sql.NullString
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CreatedAt, &path); err != nil {
			return nil, err
		}

		if path.Valid {
			post.ImagePaths, err = p.store.Image().Get(post.ID)
			if err != nil {
				return nil, err
			}
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostRepository) AddSelected(userList *[]model.User, parentID string) error {
	query := `INSERT INTO selected_users (id, user_id, parent_id) VALUES`
	var values []interface{}
	for _, user := range *userList {
		query += " (? ,? ,?),"
		id := uuid.New().String()
		values = append(values, id, user.ID, parentID)
	}

	query = strings.TrimSuffix(query, ",")

	_, err := p.store.Db.Exec(query, values...)
	return err
}

func (p *PostRepository) RemoveSelected(userList *[]model.User, parentID string) error {
	query := `DELETE FROM selected_users WHERE parent_id = ? AND (`
	values := []interface{}{parentID}
	for _, user := range *userList {
		query += " (user_id = ? AND parent_id = ?) OR"
		values = append(values, user.ID, parentID)
	}

	query = strings.TrimSuffix(query, "OR") + ")"
	_, err := p.store.Db.Exec(query, values...)

	return err
}

func (p *PostRepository) GetUserFeed(user_id string) ([]*model.Post, error) {
	query := `SELECT DISTINCT p.id, p.title, p.content, p.user_id, COALESCE(p.community_id, '') AS community_id, p.created_at, image.path
FROM post p
LEFT JOIN image ON p.id = image.parent_id
LEFT JOIN privacy pr ON pr.id = p.id
LEFT JOIN follower f ON f.source_id = ? AND f.target_id = p.user_id
LEFT JOIN selected_users su ON su.parent_id = p.id AND su.user_id = ?
LEFT JOIN member m ON m.user_id = ? AND (m.group_id = p.community_id OR m.type_id = 1)
WHERE (pr.type_id = 1 OR (pr.type_id = 3 AND su.id IS NOT NULL) OR m.id IS NOT NULL) AND f.id IS NOT NULL
ORDER BY p.created_at DESC;
`

	var posts []*model.Post
	rows, err := p.store.Db.Query(query, user_id, user_id, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var path sql.NullString
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.CommunityID, &post.CreatedAt, &path); err != nil {
			return nil, err
		}

		if path.Valid {
			post.ImagePaths, err = p.store.Image().Get(post.ID)
			if err != nil {
				return nil, err
			}
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
