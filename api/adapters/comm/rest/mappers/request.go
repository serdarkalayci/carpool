package mappers

import (
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapAddRequestRequest2Request(request dto.AddRequestRequest) (domain.Request, error) {
	var dates []time.Time
	for _, date := range request.Dates {
		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			return domain.Request{}, err
		}
		dates = append(dates, t)
	}
	return domain.Request{
		CountryID:      request.CountryID,
		Origin:         request.Origin,
		Destination:    request.Destination,
		RequestedSeats: request.RequestedSeats,
		Dates:          dates,
		State:          domain.Requested,
	}, nil
}
