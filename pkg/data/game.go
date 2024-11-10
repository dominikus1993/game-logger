package data

import "time"

type Game struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	PlayStart string    `json:"play_start"`
	PlayEnd   string    `json:"play_end"`
	Rating    int       `json:"rating"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
