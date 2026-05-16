package request

type UserRegistrationRequest struct {
	User UserRegistrationBody `json:"user"`
}

type UserRegistrationBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	User UserLoginBody `json:"user"`
}

type UserLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUserRequest struct {
	EditUserBody EditUserBody `json:"user"`
}

type EditUserBody struct {
	Image    string `json:"image"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
