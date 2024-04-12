package sqlstore

import (
	"fmt"

	"github.com/google/uuid"
)

type ImageRepository struct {
	store *Store
}

func (i *ImageRepository) Add(parentId string, paths []string) error {
	query := `INSERT INTO image (
                   id,
                   parent_id,
                   path) VALUES (?,?,?)`

	for _, path := range paths {
		_, err := i.store.Db.Exec(query,
			uuid.New().String(),
			parentId,
			path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *ImageRepository) Delete(id string) error {
	result, err := i.store.Db.Exec(`DELETE FROM image WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no image with such ID %s", id)
	}
	return nil
}

func (i *ImageRepository) DeleteAll(parentId string) error {
	result, err := i.store.Db.Exec(`DELETE FROM image WHERE parent_id = ?`, parentId)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no image with such ParentID %s", parentId)
	}
	return nil
}

func (i *ImageRepository) Update(id string, paths []string) error {
	query := `UPDATE image SET path = ? WHERE id = ?`
	var affected int64
	for _, path := range paths {
		result, err := i.store.Db.Exec(query, path, id)
		if err != nil {
			fmt.Println(err)
			return err
		}

		affected, err = result.RowsAffected()
		if err != nil {
			return err
		}
	}

	if affected == 0 {
		err := i.Add(id, paths)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *ImageRepository) Get(id string) ([]string, error) {
	query := `SELECT path FROM image WHERE parent_id = ?`
	rows, err := i.store.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var path string
	var paths []string
	for rows.Next() {
		err = rows.Scan(&path)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, nil
}
