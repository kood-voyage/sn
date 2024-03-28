package repository

import (
	"context"
	"database/sql"
	"errors"
	"social-network/followservice/model"
)

type FollowRepository interface {
	Create(ctx context.Context, sourceId, targetId string) (*model.Follow, error)
	Delete(ctx context.Context, sourceId, targetId string) error
}

type followRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) FollowRepository {
	return &followRepository{
		db: db,
	}
}

func (f *followRepository) Create(ctx context.Context, sourceId, targetId string) (*model.Follow, error) {
	query := `INSERT INTO follower (source_id, target_id)
		SELECT ?, ?
			WHERE NOT EXISTS (
    			SELECT 1
    				FROM follower
    				WHERE source_id = ? AND target_id = ?
		)`

	result, err := f.db.Exec(query, sourceId, targetId, sourceId, targetId)

	if err != nil {
		return nil, err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return nil, errors.New("already following a user")
	}

	return &model.Follow{SourceId: sourceId, TargetId: targetId}, nil
}

func (f *followRepository) Delete(ctx context.Context, sourceId, targetId string) error {
	query := `DELETE FROM follower WHERE source_id = ? AND target_id = ?`

	result, err := f.db.Exec(query, sourceId, targetId)

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
