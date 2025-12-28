package constants

import "time"

const (
	AccessPrivateKey  string        = "access"
	AccessTokenTTL    time.Duration = 15 * time.Minute
	RefreshPrivateKey string        = "refresh"
	RefreshTokenTTL   time.Duration = 7 * 24 * time.Hour
)
