package repository

import (
	"context"
	"database/sql"
	"errors"
	"social-network/requestservice/model"

	"github.com/google/uuid"
)

type RequestRepository interface {
	Create(ctx context.Context, req model.RequestReq) (*model.RequestReq, error)
	Delete(ctx context.Context, req model.RequestReq) error
	Get(ctx context.Context, req model.RequestReq) (*model.RequestReq, error)
	GetNotifications(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
	GetInvitations(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
	GetFollowrequests(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
}

type requestRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RequestRepository {
	return &requestRepository{
		db: db,
	}
}

func (r *requestRepository) Create(ctx context.Context, req model.RequestReq) (*model.RequestReq, error) {
	req.Id = uuid.New().String()
	query := `INSERT INTO request (id, type_id, source_id, target_id, parent_id, message) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, req.Id, req.TypeId, req.SourceId, req.TargetId, req.ParentId, req.Message)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *requestRepository) Delete(ctx context.Context, req model.RequestReq) error {
	query := `DELETE FROM request WHERE type_id = ? AND source_id = ? AND target_id = ?`

	result, err := r.db.Exec(query, req.TypeId, req.SourceId, req.TargetId)
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

func (r *requestRepository) Get(ctx context.Context, req model.RequestReq) (*model.RequestReq, error) {
	query := `SELECT * FROM request WHERE type_id = ? AND source_id = ? AND target_id = ?`

	var request model.RequestReq
	if err := r.db.QueryRow(query, req.TypeId, req.SourceId, req.TargetId).Scan(&request.Id, &request.TypeId, &request.SourceId, &request.TargetId, &request.ParentId, &request.Message, &request.CreatedAt); err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *requestRepository) GetNotifications(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	query := `SELECT * FROM request WHERE target_id = ? AND type_id = ?`

	var notifications []model.RequestReq

	rows, err := r.db.Query(query, user_id.Id, model.NOTIFICATION)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notification model.RequestReq
		if err := rows.Scan(&notification.Id, &notification.TypeId, &notification.SourceId, &notification.TargetId, &notification.ParentId, &notification.Message, &notification.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return &model.RequestReqs{Requests: notifications}, nil
}

func (r *requestRepository) GetInvitations(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	query := `SELECT * FROM request WHERE target_id = ? AND type_id = ?`

	var notifications []model.RequestReq

	rows, err := r.db.Query(query, user_id.Id, model.INVITE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notification model.RequestReq
		if err := rows.Scan(&notification.Id, &notification.TypeId, &notification.SourceId, &notification.TargetId, &notification.ParentId, &notification.Message, &notification.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return &model.RequestReqs{Requests: notifications}, nil
}

func (r *requestRepository) GetFollowrequests(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	query := `SELECT * FROM request WHERE target_id = ? AND type_id = ?`

	var notifications []model.RequestReq

	rows, err := r.db.Query(query, user_id.Id, model.FOLLOW)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notification model.RequestReq
		if err := rows.Scan(&notification.Id, &notification.TypeId, &notification.SourceId, &notification.TargetId, &notification.ParentId, &notification.Message, &notification.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return &model.RequestReqs{Requests: notifications}, nil
}
