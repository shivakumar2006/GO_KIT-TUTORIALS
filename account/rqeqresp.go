package account

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		ID string `json:"id"`
	}

	GetUserResponse struct {
		Email string `json:"email"`
	}
)