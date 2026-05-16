package response

type UserAuthenticationResponse struct {
	User UserAuthenticationBody `json:"user"`
}

type UserAuthenticationBody struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type UserProfileResponse struct {
	UserProfile UserProfile `json:"profile"`
}

type UserProfile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}
