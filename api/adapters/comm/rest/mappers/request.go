// Package mappers is the package that maps objects back and fort between dto and domain
package mappers

import (
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
	apierr "github.com/serdarkalayci/carpool/api/adapters/comm/rest/errors"
	"github.com/serdarkalayci/carpool/api/domain"
)

// MapAddRequestRequest2Request maps the request dto to the domain request
func MapAddRequestRequest2Request(request dto.AddRequestRequest) (domain.Request, error) {
	var dates []time.Time
	for _, date := range request.Dates {
		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			return domain.Request{}, apierr.ErrInvalidDateFormat{Date: date}
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

// MapRequest2RequestListResponse maps the domain request to the request list response dto
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

// MapRequests2RequestListResponses maps the domain requests to the request list response dtos
func MapRequests2RequestListResponses(requests *[]domain.Request) *[]dto.RequestListResponse {
	var responses []dto.RequestListResponse
	for _, request := range *requests {
		response := *MapRequest2RequestListResponse(&request)
		responses = append(responses, response)
	}
	return &responses
}

// MapRequest2RequestResponse maps the domain request to the request response dto
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
