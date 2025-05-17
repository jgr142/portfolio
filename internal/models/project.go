package models

import "time"

type Project struct {
	ID          int
	Title       string
	Description string
	Image       string
	Created     time.Time
}
