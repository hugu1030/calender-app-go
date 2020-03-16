package model

import (
	"time"
)

type Schedule struct {
	ID        int
	Title     string
	SubTitle  string
	Place     string
	FromDate  *time.Time
	ToDate    *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
