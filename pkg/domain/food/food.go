package food

import "time"

//ID represents id of Food type
type ID = uint64

// Food represents food model
type Food struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
	Location    string    `json:"location"`
	Rate        string    `json:"rate"`
	CreatedAt   time.Time `json:"created_at"`
}
