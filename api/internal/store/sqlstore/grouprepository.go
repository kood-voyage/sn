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

	if err = g.store.Image().Add(group.ID, group.ImagePaths); err != nil {
		return nil, err
	}

	return &group, nil
}

func (g *GroupRepository) Delete(groupId string) error {
	query := `DELETE FROM community WHERE id = ?`

	result, err := g.store.Db.Exec(query, groupId)
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
	if err = g.store.Privacy().Delete(groupId); err != nil {
		return err
	}

	if err = g.store.Image().DeleteAll(groupId); err != nil {
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

	if err = g.store.Privacy().Update(group.ID, privacy); err != nil {
		return err
	}

	if err = g.store.Image().Update(group.ID, group.ImagePaths); err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) Get(groupId string) (*model.Group, error) {
	query := `SELECT * FROM community WHERE name = ?`
	var group model.Group
	if err := g.store.Db.QueryRow(query, groupId).Scan(
		&group.ID,
		&group.CreatorID,
		&group.Name,
		&group.Description,
	); err != nil {
		return nil, err
	}

	gMembers, err := g.Members(group.ID)
	if err != nil {
		return nil, err
	}
	group.Members = *gMembers
	paths, err := g.store.Image().Get(group.ID)
	if err != nil {
		return nil, err
	}
	group.ImagePaths = paths

	return &group, nil
}

func (g *GroupRepository) Members(groupId string) (*[]model.User, error) {
	query := `SELECT member.user_id, member.type_id 
	FROM member
	JOIN member_type on  member.type_id = member_type.id
	WHERE group_id = ?`

	var users []model.User
	rows, err := g.store.Db.Query(query, groupId)
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

func (g *GroupRepository) IsMember(groupId, userId string) (bool, error) {
	query := `SELECT user_id FROM member WHERE group_id = ? AND user_id = ?`

	var user model.User
	if err := g.store.Db.QueryRow(query, groupId, userId).Scan(&user.ID); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (g *GroupRepository) AddMember(groupId, userId string) error {
	query := `INSERT INTO member (id, user_id, group_id, type_id) VALUES (?, ?, ?, ?)`

	_, err := g.store.Db.Exec(query, uuid.New().String(), userId, groupId, 1)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) GetAll(types model.Type) (*[]model.Group, error) {
	query := `SELECT * FROM community`

	var groups []model.Group
	rows, err := g.store.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var group model.Group
		if err := rows.Scan(&group.ID, &group.CreatorID, &group.Name, &group.Description); err != nil {
			return nil, err
		}
		privacy, err := g.store.Privacy().Check(group.ID) 
		if err != nil {
			return nil, err
		}
		group.Privacy, err = types.IntToString(privacy)
		if err != nil {
			return nil, err
		}

		group.ImagePaths, err = g.store.Image().Get(group.ID)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	return &groups, nil
}
