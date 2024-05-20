package sqlstore

import (
	"errors"
	"social-network/internal/model"

	"github.com/google/uuid"
)

type RequestRepository struct {
	store *Store
}

func (r *RequestRepository) Create(request model.Request) (*model.Request, error) {
	request.ID = uuid.New().String()
	query := `INSERT INTO request (id, type_id, source_id, target_id, parent_id, message) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.store.Db.Exec(query, request.ID, request.TypeID, request.SourceID, request.TargetID, request.ParentID, request.Message)
	if err != nil {
		return nil, err
	}

	sourcUser, err := r.store.User().Get(request.SourceID)
	if err != nil {
		return nil, err
	}
	sourcUser.Sanitize()

	targetUser, err := r.store.User().Get(request.TargetID)
	if err != nil {
		return nil, err
	}
	targetUser.Sanitize()

	request.SourceInformation = *sourcUser
	request.TargetInformation = *targetUser


	return &request, nil
}

func (r *RequestRepository) Delete(request model.Request) error {
	query := `DELETE FROM request WHERE type_id = ? AND source_id = ? AND target_id = ?`

	result, err := r.store.Db.Exec(query, request.TypeID, request.SourceID, request.TargetID)
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

func (r *RequestRepository) DeleteByID(id string) error {
	query := `DELETE FROM request WHERE id = ?`

	_, err := r.store.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RequestRepository) Get(request model.Request) (*model.Request, error) {
	query := `SELECT * FROM request WHERE type_id = ? AND source_id = ? AND target_id = ?`

	var req model.Request
	if err := r.store.Db.QueryRow(query, request.TypeID, request.SourceID, request.TargetID).Scan(&req.ID, &req.TypeID, &req.SourceID, &req.TargetID, &req.ParentID, &req.Message, &req.CreatedAt); err != nil {
		return nil, err
	}

	return &req, nil
}


func (r *RequestRepository) GetGroups(request model.Request) (*model.Request, error) {
	query := `SELECT * FROM request WHERE type_id = ? AND source_id = ? AND target_id = ? AND parent_id = ?`

	var req model.Request
	if err := r.store.Db.QueryRow(query, request.TypeID, request.SourceID, request.TargetID, request.ParentID).Scan(&req.ID, &req.TypeID, &req.SourceID, &req.TargetID, &req.ParentID, &req.Message, &req.CreatedAt); err != nil {
		return nil, err
	}

	return &req, nil
}

