package model

type Group struct {
	ID          string   `db:"id" json:"id" validate:"required"`
	CreatorID   string   `db:"creator_id" json:"creator_id"`
	Name        string   `db:"name" json:"name" validate:"required|min_len:2|max_len:25"`
	Description string   `db:"description" json:"description" validate:"required"`
	ImagePaths  []string `db:"path" json:"image_path"`
	Privacy     string   `json:"privacy" validate:"required|contains:public,private"`
	Members     []User   `json:"members"`
}

func NewGroup() *Group {
	return &Group{}
}
