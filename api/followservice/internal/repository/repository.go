package repository

import (
	"context"
	"fmt"
)

type FollowRepository interface {
	Create(ctx context.Context, sourceId, targetId string) error
	Delete(ctx context.Context, sourceId, targetId string) error
}

type followRepository struct {
	// db *sql.DB
}

func NewRepository() FollowRepository {
	return &followRepository{
		// db: db,
	}
}

func (f *followRepository) Create(ctx context.Context, sourceId, targetId string) error {
	fmt.Println("Inserting something to db", sourceId, targetId)
	return nil
}

func (f *followRepository) Delete(ctx context.Context, sourceId, targetId string) error {
	fmt.Println("Deleting somethng from db", sourceId, targetId)
	return nil
}
