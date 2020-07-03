package ac

// RegisterAndConfirmResponse response contract
type RegisterAndConfirmResponse struct {
	ID      int
	Name    string
	Email   string
	JWT     string
	Success bool
}

// CheckUserAuthResponse response contract
type CheckUserAuthResponse struct {
	ID      int
	Name    string
	Email   string
	JWT     string
	Success bool
}
