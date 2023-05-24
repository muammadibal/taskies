package user

type UserService interface {
	GetUserByID(id string) (*User, error)
}

type userService struct {
	userService UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) GetUserByID(id string) (*User, error) {
	// Add any business logic or validation here if needed
	return s.userService.GetByID(id)
}
