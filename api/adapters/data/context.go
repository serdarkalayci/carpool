package data

import "github.com/serdarkalayci/carpool/api/application"

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	UserRepository      application.UserRepository
	HealthRepository    application.HealthRepository
	GeographyRepository application.GeographyRepository
	TripRepository      application.TripRepository
}
