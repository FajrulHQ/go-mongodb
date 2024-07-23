package project

import (
	"time"
)

type Projects struct {
	ID      string     `json:"_id"`
	Name    string     `json:"name"`
	Desc    string     `json:"description"`
	Link    string     `json:"link"`
	Start   *time.Time `json:"start,omitempty"`
	End     *time.Time `json:"end,omitempty"`
	Created time.Time  `json:"created"`
	Updated time.Time  `json:"updated"`
	Image   string     `json:"image"`
}
