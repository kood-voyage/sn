package repository

import (
	"context"
	"database/sql"
	"errors"
	"social-network/privacyservice/model"
)

type PrivacyRepository interface {
	Create(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error)
	Update(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error)
	Delete(ctx context.Context, parent_id string) error
	Get(ctx context.Context, parent_id string) (*model.PrivacyReq, error)
}

type privacyRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) PrivacyRepository {
	return &privacyRepository{
		db: db,
	}
}

func (p *privacyRepository) Create(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error) {
	query := `INSERT INTO privacy (id, type_id) VALUES (?,?)`

	_, err := p.db.Exec(query, parent_id, privacy)
	if err != nil {
		return nil, err
	}
	return &model.PrivacyReq{
		ParentId: parent_id,
		Privacy:  privacy,
	}, nil
}

func (p *privacyRepository) Update(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error) {
	query := `UPDATE privacy SET type_id = ? WHERE id = ?`

	result, err := p.db.Exec(query, privacy, parent_id)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("PRIVACY: no rows affected")
	}

	return &model.PrivacyReq{
		ParentId: parent_id,
		Privacy:  privacy,
	}, nil
}

func (p *privacyRepository) Delete(ctx context.Context, parent_id string) error {
	query := `DELETE FROM privacy WHERE id = ?`

	result, err := p.db.Exec(query, parent_id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (p *privacyRepository) Get(ctx context.Context, parent_id string) (*model.PrivacyReq, error) {
	query := `SELECT * FROM privacy WHERE id = ?`

	var privacy model.PrivacyReq
	if err := p.db.QueryRow(query, parent_id).Scan(&privacy.ParentId, &privacy.Privacy); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("privacy entry does not exist")
		}
		return nil, err
	}
	return &privacy, nil
}
