package model

import "time"

type Game struct {
	Id           string        `json:"id"`
	Title        string        `json:"title"`
	Playthroughs []Playthrough `json:"playthroughs"`
}

type Playthrough struct {
	StartDate   time.Time  `json:"start_date"`
	FinishDate  *time.Time `json:"finish_date,omitempty"`
	Platform    string     `json:"platform,omitempty"`
	HoursPlayed *int       `json:"hours_played,omitempty"`
	Rating      *int       `json:"rating,omitempty"`
	Notes       string     `json:"notes,omitempty"`
}

func (g *Game) AddPlaythrough(pt Playthrough) {
	g.Playthroughs = append(g.Playthroughs, pt)
}
