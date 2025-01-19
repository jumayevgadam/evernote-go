package jwt

// Tokens model keeps access token and refresh token.
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token	"`
}
