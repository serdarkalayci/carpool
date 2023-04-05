// Package application is the package that holds the application logic between database and communication layers
package application

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	UserRepository         UserRepository
	HealthRepository       HealthRepository
	GeographyRepository    GeographyRepository
	TripRepository         TripRepository
	ConversationRepository ConversationRepository
	RequestRepository      RequestRepository
}
