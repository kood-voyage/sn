package store

import "social-network/internal/model"

type UserRepository interface {
	Create(user model.User, privacy int) (*model.User, error)
	GetFollowers(userID string) ([]model.User, error)
	GetFollowing(userID string) ([]model.User, error)
	IsFollowing(source_id, target_id string) (bool, error)
	GetNotifications(user_id string) ([]model.Request, error)
	GetAll() ([]model.User, error)
	SetDescription(user_id, description string) error
	SetCover(user_id, cover string) error
	SetAvatar(user_id, avatar string) error
	Get(user_id string) (*model.User, error)
}

type SessionRepository interface {
	Check(access_id string) (*model.Session, error)
	Delete(access_id string) error
	Update(access_id string, session model.Session) (*model.Session, error)
	Create(session model.Session) (*model.Session, error)
	DeleteByUser(user_id string) error
	CheckByUserId(user_id string) (*model.Session, error)
}

type FollowRepository interface {
	Create(follower model.Follower) error
	Delete(follower model.Follower) error
}

type RequestRepository interface {
	Create(request model.Request) (*model.Request, error)
	Delete(request model.Request) error
	Get(request model.Request) (*model.Request, error)
	DeleteByID(id string) error
	GetGroups(request model.Request) (*model.Request, error)
}

type PostRepository interface {
	Create(post *model.Post, privacy int) error
	Delete(id string) error
	Get(id string) (*model.Post, error)
	Update(post *model.Post, privacy int) error
	GetUsers(source_id, target_id string) ([]model.Post, error)
	AddSelected(userList *[]model.User, parentID string) error
	RemoveSelected(userList *[]model.User, parentID string) error
	GetUserFeed(user_id string) ([]*model.Post, error)
}

type CommentRepository interface {
	Create(post *model.Comment) error
	Delete(postID, userID string) error
	// GetAll returns all comments to single post
	GetAll(id string) (*[]model.Comment, error)
	IsAuthor(comment *model.Comment, userId string) bool
	Update(comment *model.Comment) error
}

type GroupRepository interface {
	Create(group model.Group, privacy int) (*model.Group, error)
	Delete(groupId string) error
	Update(group model.Group, privacy int) error
	Get(groupId string) (*model.Group, error)
	Members(groupId string) (*[]model.User, error)
	IsMember(groupId, userId string) (bool, error)
	AddMember(groupId, userId string) error
	GetAll(model.Type) (*[]model.Group, error)
	GetPosts(group_id string) ([]*model.Post, error)
	GetAllEvents(group_id, source_id string) ([]*model.Event, error)
	GetInvitedUsers(group_id string) ([]string, error)
}

type PrivacyRepository interface {
	Set(parent_id string, privacy int) error
	Update(parent_id string, privacy int) error
	Delete(parent_id string) error
	Check(parent_id string) (int, error)
}

type EventRepository interface {
	Create(event *model.Event) error
	Update(event *model.Event) error
	Delete(eventId string) error
	Get(eventId string) (*model.Event, error)
	Register(userid, eventId string, opt int) error
	AllParticipants(eventId string) ([]*model.User, error)
}

type ImageRepository interface {
	Add(parentId string, paths []string) error
	Delete(id string) error
	DeleteAll(parentId string) error
	Update(parentId string, paths []string) error
	Get(parentId string) ([]string, error)
}

type ChatRepository interface {
	Create(chat model.Chat) (*model.Chat, error)
	AddUser(user model.User, chat model.Chat) error
	GetUsers(chat model.Chat) ([]model.User, error)
	AddLine(line *model.ChatLine) (*model.ChatLine, error)
	GetAll(user_id string) ([]*model.Chat, error)
	Load(chat_id, user_id string) ([]*model.ChatLine, error)
	GetChatsForUser(userID string) ([]*model.User, error)
	// TO-DO
	// DeleteUser()
}
