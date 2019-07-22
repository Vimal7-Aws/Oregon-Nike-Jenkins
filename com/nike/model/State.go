package model

import "time"

type State struct {
	Name    string    `json:"name"`
	Rivers  []string  `json:"rivers"`
	Created time.Time `json:"time"`
}
