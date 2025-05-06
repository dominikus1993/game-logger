package model

type Game struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	StartDate   string  `json:"start_date"`
	FinishDate  *string `json:"finish_date,omitempty"`
	Platform    *string `json:"platform,omitempty"`
	HoursPlayed *uint16 `json:"hours_played,omitempty"`
	Rating      *uint16 `json:"rating,omitempty"`
	Notes       *string `json:"notes,omitempty"`
}
