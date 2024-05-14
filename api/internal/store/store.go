package store

type Store interface {
	User() UserRepository
	Session() SessionRepository
	Follow() FollowRepository
	Request() RequestRepository
	Post() PostRepository
	Group() GroupRepository
	Comment() CommentRepository
	Privacy() PrivacyRepository
	Event() EventRepository
	Chat() ChatRepository
	Image() ImageRepository
}
