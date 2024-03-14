package sqlstore

import (
	"fmt"
	"social-network/internal/model"
)

type EventRepository struct {
	store *Store
}

func (e EventRepository) Create(event *model.Event) error {
	query := `INSERT INTO event (
                   id,
                   user_id,
                   group_id,
                   name,
                   description,
                   date) VALUES (?,?,?,?,?,?)`

	_, err := e.store.Db.Exec(query,
		event.ID,
		event.UserID,
		event.GroupID,
		event.Name,
		event.Description,
		event.Date)
	if err != nil {
		return err
	}

	return nil
}

func (e EventRepository) Update(event *model.Event) error {
	query := `UPDATE event SET name = ?, description = ?, date = ? WHERE id = ?`

	_, err := e.store.Db.Exec(query, event.Name, event.Description, event.Date, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e EventRepository) Delete(eventId string) error {
	query := `DELETE FROM event WHERE id = ?`

	result, err := e.store.Db.Exec(query, eventId)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no event with such ID %s", eventId)
	}
	return nil
}

func (e EventRepository) Get(eventId string) (*model.Event, error) {
	event := &model.Event{}
	query := `SELECT * FROM event WHERE id = ?`

	err := e.store.Db.QueryRow(query, eventId).Scan(
		&event.ID,
		&event.UserID,
		&event.GroupID,
		&event.Name,
		&event.Description,
		&event.CreatedAt,
		&event.Date)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e EventRepository) Register(eventId, opt string) error {
	//TODO implement me
	panic("implement me")
}
