package sqlstore

import (
	models "social-network/internal/model"

	"github.com/google/uuid"
)

type RequestRepository struct {
	store *Store
}

func (r *RequestRepository) Create(request models.Request) error {
	request.ID = uuid.New().String()
	query := `INSERT INTO request (id, type_id, source_id, target_id, message) VALUES (?, ?, ?, ?, ?)`

	_, err := r.store.Db.Exec(query, request.ID, request.TypeID, request.SourceID, request.TargetID, request.Message)
	if err != nil {
		return err
	}

	return nil
}