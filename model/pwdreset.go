package model

import (
	"time"
)

type PwdReset struct {
	Email     string
	Token     string
	CreatedAt time.Time
}
