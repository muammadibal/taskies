package user

import "github.com/gofiber/fiber/v2"

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	// Get the ID from the request paramfiber
	id := c.Params("id")
	// Parse the ID to uint or handle the error
	// ...

	// Call the UserServuserService method
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		// Handle the error
		// ...
	}

	// Return the user as JSON response
	return c.JSON(user)
}
