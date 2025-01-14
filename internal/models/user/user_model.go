package user

// we will separate user models inside this package.
// dto and dao models.

// SignUpReq model for registering to application.
type SignUpReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	Email    string `json:"email" binding:"required,email"`
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

// LoginReq model.
type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AllUserData model contains needed all fields for user.
type AllUserData struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}
