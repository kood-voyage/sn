package sqlstore

import (
	"database/sql"
	"fmt"
	"social-network/internal/model"

	"github.com/google/uuid"
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
	query := `
	SELECT e.id, e.user_id, e.group_id, e.name, e.description, e.created_at, e.date,
		   u.id, u.username, u.email, u.timestamp, u.date_of_birth, u.first_name, u.last_name, u.description, u.avatar, u.cover
	FROM event e
	JOIN user u ON e.user_id = u.id
	WHERE e.id = ?
`

	err := e.store.Db.QueryRow(query, eventId).Scan(
		&event.ID,
		&event.UserID,
		&event.GroupID,
		&event.Name,
		&event.Description,
		&event.CreatedAt,
		&event.Date,
		&event.UserInformation.ID,
		&event.UserInformation.Username,
		&event.UserInformation.Email,
		&event.UserInformation.CreatedAt,
		&event.UserInformation.DateOfBirth,
		&event.UserInformation.FirstName,
		&event.UserInformation.LastName,
		&event.UserInformation.Description,
		&event.UserInformation.Avatar,
		&event.UserInformation.Cover,
	)
	if err != nil {
		return nil, err
	}

	participants, err := e.AllParticipants(eventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return event, nil
		}
		return nil, err
	}

	event.Participants = participants

	return event, nil
}

func (e *EventRepository) AllParticipants(eventId string) ([]*model.User, error) {
	participantsQuery := `
	SELECT u.id, u.username, u.email, u.timestamp, u.date_of_birth, u.first_name, u.last_name, u.description, u.avatar, u.cover, eot.description
	FROM event_registered_users eru
	JOIN user u ON eru.user_id = u.id
	JOIN event_option_type eot ON eru.type_id = eot.id 
	WHERE eru.event_id = ?
`

	rows, err := e.store.Db.Query(participantsQuery, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []*model.User
	for rows.Next() {
		var participant model.User
		if err := rows.Scan(
			&participant.ID,
			&participant.Username,
			&participant.Email,
			&participant.CreatedAt,
			&participant.DateOfBirth,
			&participant.FirstName,
			&participant.LastName,
			&participant.Description,
			&participant.Avatar,
			&participant.Cover,
			&participant.EventStatus,
		); err != nil {
			return nil, err
		}
		participants = append(participants, &participant)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return participants, nil
}

func (e EventRepository) Register(userId, eventId string, opt int) error {
	query := `
        INSERT INTO event_registered_users (id, type_id, user_id, event_id)
        SELECT ?, ?, ?, ?
        WHERE NOT EXISTS (
            SELECT 1 FROM event_registered_users WHERE user_id = ? AND event_id = ?
        )
    `

	_, err := e.store.Db.Exec(query, uuid.New().String(), opt, userId, eventId, userId, eventId)
	if err != nil {
		return err
	}

	return nil
}
