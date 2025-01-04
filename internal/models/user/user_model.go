package user

// we will separate user models inside this package.
// dto and dao models.

// SignUpReq model for registering to application.
type SignUpReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

// SignUpReqData model represent db model.
type SignUpReqData struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

// ToPsqlDBStorage func sends request to storage.
func (s *SignUpReq) ToPsqlDBStorage() SignUpReqData {
	return SignUpReqData{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}
}
