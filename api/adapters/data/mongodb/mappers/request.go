// Package mappers is the package that maps objects back and fort between dao and domain
package mappers

import (
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapRequest2RequestDAO maps a domain.Request to a dao.RequestDAO
func MapRequest2RequestDAO(request *domain.Request) *dao.RequestDAO {
	id, _ := primitive.ObjectIDFromHex(request.ID)
	requesterID, _ := primitive.ObjectIDFromHex(request.RequesterID)
	countryID, _ := primitive.ObjectIDFromHex(request.CountryID)
	var dates []primitive.DateTime
	for _, date := range request.Dates {
		dates = append(dates, primitive.NewDateTimeFromTime(date))
	}
	return &dao.RequestDAO{
		ID:             id,
		RequesterID:    requesterID,
		RequesterName:  request.RequesterName,
		CountryID:      countryID,
		Origin:         request.Origin,
		Destination:    request.Destination,
		Dates:          dates,
		RequestedSeats: request.RequestedSeats,
		State:          int(request.State),
	}
}

// MapRequestDAO2Request maps a dao.RequestDAO to a domain.Request
func MapRequestDAO2Request(requestDAO *dao.RequestDAO) *domain.Request {
	var dates []time.Time
	for _, date := range requestDAO.Dates {
		dates = append(dates, date.Time())
	}
	return &domain.Request{
		ID:             requestDAO.ID.Hex(),
		RequesterID:    requestDAO.RequesterID.Hex(),
		RequesterName:  requestDAO.RequesterName,
		CountryID:      requestDAO.CountryID.Hex(),
		Origin:         requestDAO.Origin,
		Destination:    requestDAO.Destination,
		RequestedSeats: requestDAO.RequestedSeats,
		Dates:          dates,
		State:          domain.RequestState(requestDAO.State),
	}
}

// MapRequesDAOs2Requests maps a slice of dao.RequestDAO to a slice of domain.Request
func MapRequesDAOs2Requests(requestDAOs []dao.RequestDAO) *[]domain.Request {
	var requests []domain.Request
	for _, requestDAO := range requestDAOs {
		req := *MapRequestDAO2Request(&requestDAO)
		requests = append(requests, req)
	}
	return &requests
}
