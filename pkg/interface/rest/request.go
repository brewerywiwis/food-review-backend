package rest

type CreateFoodRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Location    string `json:"location"`
	Rate        string `json:"rate"`
}
