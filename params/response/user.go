package response

type UserAuthenticationResponse struct {
	User UserAuthenticationBody `json:"user"`
}

type UserAuthenticationBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Username string `json:"username"`
	Bio string `json:"bio"`
	Image string `json:"image"`
}

