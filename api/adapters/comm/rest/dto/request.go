package dto

type AddRequestRequest struct {
	CountryID      string   `json:"countryid" validate:"required"`
	Origin         string   `json:"origin" validate:"required"`
	Destination    string   `json:"destination" validate:"required"`
	RequestedSeats int      `json:"requestedseats" validate:"required"`
	Dates          []string `json:"dates" validate:"required"`
}

type RequestListResponse struct {
	ID             string   `json:"id"`
	Origin         string   `json:"origin"`
	Destination    string   `json:"destination"`
	RequestedSeats int      `json:"requestedseats"`
	Dates          []string `json:"dates"`
	State          string
}

type RequestResponse struct {
	ID             string   `json:"id"`
	RequesterID    string   `json:"requesterid"`
	RequesterName  string   `json:"requestername"`
	CountryID      string   `json:"countryid"`
	Origin         string   `json:"origin"`
	Destination    string   `json:"destination"`
	RequestedSeats int      `json:"requestedseats"`
	Dates          []string `json:"dates"`
	State          string   `json:"state"`
}
