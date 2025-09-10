package domain

import "time"

type Twb struct {
	ID        string        `json:"id"`
	Path      string        `json:"path"`
	ExpiresIn time.Duration `json:"expires_in"`
}
