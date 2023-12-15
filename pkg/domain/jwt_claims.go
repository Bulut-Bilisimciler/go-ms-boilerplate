package domain

import "time"

// JWTClaims from user microservice
type JWTClaims struct {
	AuthId     int64      `json:"auth_id"`
	Name       string     `json:"name"`
	Slug       string     `json:"slug"`
	Role       string     `json:"role"`
	Status     int        `json:"status"`
	VerifiedAt *time.Time `json:"verified_at"`
}
