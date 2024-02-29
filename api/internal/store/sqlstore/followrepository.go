package sqlstore

import (
	models "social-network/internal/model"

	"github.com/google/uuid"
)

type FollowRepository struct {
	store *Store
}

func (f *FollowRepository) Create(follower models.Follower) error {
	follower.ID = uuid.New().String()
	query := `INSERT INTO follower (id, source_id, target_id) VALUES (?, ?, ?)`

	_, err := f.store.Db.Exec(query, follower.ID, follower.SourceID, follower.TargetID)

	if err != nil {
		return err
	}

	return nil
}

func (f *FollowRepository) Delete(follower models.Follower) error {
	query := `DELETE FROM follower WHERE source_id = ? AND target_id = ?`

	_, err := f.store.Db.Exec(query, follower.SourceID, follower.TargetID)

	if err != nil {
		return err
	}

	return nil
}
