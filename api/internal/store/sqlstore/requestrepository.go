package sqlstore

import (
	"social-network/internal/model"

	"github.com/google/uuid"
)

type RequestRepository struct {
	store *Store
}

func (r *RequestRepository) Create(request model.Request) error {
	request.ID = uuid.New().String()
	query := `INSERT INTO request (id, type_id, source_id, target_id, message) VALUES (?, ?, ?, ?, ?)`

	_, err := r.store.Db.Exec(query, request.ID, request.TypeID, request.SourceID, request.TargetID, request.Message)
	if err != nil {
		return err
	}

	return nil
}

func (r *RequestRepository) Delete(request model.Request) error {
	query := `DELETE FROM request WHERE type_id = ? AND source_id = ? AND target_id = ?`

	_, err := r.store.Db.Exec(query, request.TypeID, request.SourceID, request.TargetID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RequestRepository) Get(request model.Request) (*model.Request, error) {
	query := `SELECT * FROM request WHERE type_id = ? AND source_id = ? AND target_id = ?`

	var req model.Request
	if err := r.store.Db.QueryRow(query, request.TypeID, request.SourceID, request.TargetID).Scan(&req.ID, &req.TypeID, &req.SourceID, &req.TargetID, &req.Message, &req.CreatedAt); err != nil {
		return nil, err
	}

	return &req, nil
}
