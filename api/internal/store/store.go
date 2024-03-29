package store

type Store interface {
	User() UserRepository
	Follow() FollowRepository
	Request() RequestRepository
	Post() PostRepository
	Group() GroupRepository
	Comment() CommentRepository
	Privacy() PrivacyRepository
	Event() EventRepository
	Image() ImageRepository
}
