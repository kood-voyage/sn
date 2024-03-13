package sqlstore

import "social-network/internal/model"

type EventRepository struct {
	store *Store
}

func (e EventRepository) Create(event *model.Event) (*model.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) Update(event *model.Event) error {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) Delete(eventId string) error {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) Get(eventId string) error {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) Register(eventId, opt string) error {
	//TODO implement me
	panic("implement me")
}
