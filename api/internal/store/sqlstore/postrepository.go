package sqlstore

import "social-network/internal/model"

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

	return nil
}

//func (p PostRepository) Update() {
//
//}

func (p PostRepository) Get(id string) (*model.Post, error) {
	post := &model.Post{}

	return post, nil
}
