package shared

type CreateUserArgs struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CredentialArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
