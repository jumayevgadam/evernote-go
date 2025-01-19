package constants

import "time"

const (
	// ShutdownTimeOut is.
	ShutdownTimeOut = 5 * time.Second
)

// JwtTokenExpiry times.
const (
	// AccessTokenExpiryTime is.
	AccessTokenExpiryTime = 5 * time.Second

	// RefreshTokenExpiryTime is.
	RefreshTokenExpiryTime = 24 * time.Hour
)
