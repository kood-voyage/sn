package sqlstore

import (
	"errors"
	"social-network/internal/model"

	"github.com/google/uuid"
)

type FollowRepository struct {
	store *Store
}

func (f *FollowRepository) Create(follower model.Follower) error {
	follower.ID = uuid.New().String()
	query := `INSERT INTO follower (id, source_id, target_id)
SELECT ?, ?, ?
WHERE NOT EXISTS (
    SELECT 1
    FROM follower
    WHERE source_id = ? AND target_id = ?
)`

	result, err := f.store.Db.Exec(query, follower.ID, follower.SourceID, follower.TargetID, follower.SourceID, follower.TargetID)

	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("already following a user")
	}

	return nil
}

func (f *FollowRepository) Delete(follower model.Follower) error {
	query := `DELETE FROM follower WHERE source_id = ? AND target_id = ?`

	result, err := f.store.Db.Exec(query, follower.SourceID, follower.TargetID)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}
