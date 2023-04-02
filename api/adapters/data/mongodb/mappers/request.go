package mappers

import (
	"time"

	"github.com/serdarkalayci/carpool/api/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/carpool/api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func MapRequestDAO2Request(requestDAO *dao.RequestDAO) *domain.Request {
	var dates []time.Time
	for _, date := range requestDAO.Dates {
		dates = append(dates, date.Time())
	}
	return &domain.Request{
		ID:            requestDAO.ID.Hex(),
		RequesterID:   requestDAO.RequesterID.Hex(),
		RequesterName: requestDAO.RequesterName,
		CountryID:     requestDAO.CountryID.Hex(),
		Origin:        requestDAO.Origin,
		Destination:   requestDAO.Destination,
		Dates:         dates,
		State:         domain.RequestState(requestDAO.State),
	}
}

func MapRequesDAOs2Requests(requestDAOs []dao.RequestDAO) *[]domain.Request {
	var requests *[]domain.Request
	for _, requestDAO := range requestDAOs {
		*requests = append(*requests, *MapRequestDAO2Request(&requestDAO))
	}
	return requests
}
