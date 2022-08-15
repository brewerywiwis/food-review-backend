package food

import "time"

// Food represents food model
type Food struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
	Location    string    `json:"location"`
	Rate        string    `json:"rate"`
	CreatedAt   time.Time `json:"created_at"`
}
