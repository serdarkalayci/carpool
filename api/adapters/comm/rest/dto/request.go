package dto

type AddRequestRequest struct {
	CountryID      string   `json:"countryid" validate:"required"`
	Origin         string   `json:"origin" validate:"required"`
	Destination    string   `json:"destination" validate:"required"`
	RequestedSeats int      `json:"requestedseats" validate:"required"`
	Dates          []string `json:"dates" validate:"required"`
}
