package ac

// RegisterAndConfirmRequest request contract
type RegisterAndConfirmRequest struct {
	Name     string
	Email    string
	Password string
	Consent  bool
}

// CheckUserAuthRequest request contract
type CheckUserAuthRequest struct {
	JWT string
}
