package request

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type SendEmailRequest struct {
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Receiver string `json:"receiver"`
}
