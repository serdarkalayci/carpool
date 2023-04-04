package mappers

import (
	"fmt"
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	"github.com/serdarkalayci/carpool/api/domain"
)

func MapAddRequestRequest2Request(request dto.AddRequestRequest) (domain.Request, error) {
	var dates []time.Time
	for _, date := range request.Dates {
		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			return domain.Request{}, errInvalidDateFormat{date: date}
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

func MapRequest2RequestListResponse(request *domain.Request) *dto.RequestListResponse {
	var dates []string
	for _, date := range request.Dates {
		dates = append(dates, date.Format("2006-01-02"))
	}
	return &dto.RequestListResponse{
		ID:             request.ID,
		Origin:         request.Origin,
		Destination:    request.Destination,
		RequestedSeats: request.RequestedSeats,
		Dates:          dates,
		State:          request.State.String(),
	}
}

func MapRequests2RequestListResponses(requests *[]domain.Request) *[]dto.RequestListResponse {
	var responses []dto.RequestListResponse
	for _, request := range *requests {
		response := *MapRequest2RequestListResponse(&request)
		responses = append(responses, response)
	}
	return &responses
}

func MapRequest2RequestResponse(request *domain.Request) *dto.RequestResponse {
	var dates []string
	for _, date := range request.Dates {
		dates = append(dates, date.Format("2006-01-02"))
	}
	return &dto.RequestResponse{
		ID:             request.ID,
		RequesterID:    request.RequesterID,
		RequesterName:  request.RequesterName,
		CountryID:      request.CountryID,
		Origin:         request.Origin,
		Destination:    request.Destination,
		RequestedSeats: request.RequestedSeats,
		Dates:          dates,
		State:          request.State.String(),
	}
}

type errInvalidDateFormat struct {
	date string
}

func (e errInvalidDateFormat) Error() string {
	return fmt.Sprintf("invalid date %s for the format: yyyy-MM-dd", e.date)
}
