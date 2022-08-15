package food

import "time"

// Food represents food model
type Food struct {
	ID          int64
	Name        string
	Description string
	Img         string
	Location    string
	Rate        string
	CreatedAt   time.Time
}
