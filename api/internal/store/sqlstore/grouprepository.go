package sqlstore

import (
	"database/sql"
	"errors"
	"social-network/internal/model"

	"github.com/google/uuid"
)

type GroupRepository struct {
	store *Store
}

func (g *GroupRepository) Create(group model.Group, privacy int) (*model.Group, error) {
	query := `INSERT INTO community (id, creator_id, name, description) VALUES (?, ?, ?, ?)`

	_, err := g.store.Db.Exec(query, group.ID, group.CreatorID, group.Name, group.Description)
	if err != nil {
		return nil, err
	}

	if err := g.store.Privacy().Set(group.ID, privacy); err != nil {
		return nil, err
	}

	return &group, nil
}

func (g *GroupRepository) Delete(group_id string) error {
	query := `DELETE FROM community WHERE id = ?`

	result, err := g.store.Db.Exec(query, group_id)
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

	//remove from privacy table
	if err := g.store.Privacy().Delete(group_id); err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) Update(group model.Group, privacy int) error {
	query := `UPDATE community SET name = ?, description = ? WHERE id = ?`

	_, err := g.store.Db.Exec(query, group.Name, group.Description, group.ID)
	if err != nil {
		return err
	}

	if err := g.store.Privacy().Update(group.ID, privacy); err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) Get(group_id string) (*model.Group, error) {
	query := `SELECT * FROM community WHERE id = ?`

	var group model.Group
	if err := g.store.Db.QueryRow(query, group_id).Scan(&group.ID, &group.CreatorID, &group.Name, &group.Description); err != nil {
		return nil, err
	}

	g_members, err := g.Members(group_id)
	if err != nil {
		return nil, err
	}
	group.Members = *g_members

	return &group, nil
}

func (g *GroupRepository) Members(group_id string) (*[]model.User, error) {
	query := `SELECT member.user_id, member.type_id 
	FROM member
	JOIN member_type on  member.type_id = member_type.id
	WHERE group_id = ?`

	var users []model.User
	rows, err := g.store.Db.Query(query, group_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.MemberType); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (g *GroupRepository) IsMember(group_id, user_id string) error {
	query := `SELECT user_id FROM member WHERE group_id = ? AND user_id = ?`

	var user model.User
	if err := g.store.Db.QueryRow(query, group_id, user_id).Scan(&user.ID); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user is not a group member")
		}
	}

	return nil
}

func (g *GroupRepository) AddMember(group_id, user_id string) error {
	query := `INSERT INTO member id, user_id, group_id, type_id VALUES (?, ?, ?, ?)`

	_, err := g.store.Db.Exec(query, uuid.New().String(), user_id, group_id, 1)
	if err != nil {
		return err
	}

	return nil
}
