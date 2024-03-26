package repository

import "fmt"

type FollowRepository interface {
	Create(sourceId, targetId string) error
	Delete(sourceId, targetId string) error
}

type followRepository struct {
	// db *sql.DB
}

func NewRepository() *followRepository {
	return &followRepository{
		// db: db,
	}
}

func (f *followRepository) Create(sourceId, targetId string) error {
	fmt.Println("Inserting something to db", sourceId, targetId)
	return nil
}

func (f *followRepository) Delete(sourceId, targetId string) error {
	fmt.Println("Deleting somethng from db", sourceId, targetId)
	return nil
}
