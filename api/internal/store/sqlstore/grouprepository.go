package sqlstore

import "social-network/internal/model"

type GroupRepository struct {
	store *Store
}

func (g *GroupRepository) Create(group model.Group) (*model.Group, error) {
	query := `INSERT INTO community (id, creator_id, name, description) VALUES (?, ?, ?, ?)`

	_, err := g.store.Db.Exec(query, group.ID, group.CreatorID, group.Name, group.Description)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (g *GroupRepository) Delete(group_id string) error {
	query := `DELETE FROM community WHERE id = ?`

	_, err := g.store.Db.Exec(query, group_id)
	if err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) Update(group model.Group) error {
	query := `UPDATE community SET name = ?, description = ? WHERE id = ?`

	_, err := g.store.Db.Exec(query, group.Name, group.Description, group.ID)
	if err != nil {
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

	return &group, nil
}