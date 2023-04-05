// Package application is the package that holds the application logic between database and communication layers
package application

// DataContextCarrier is the interface to be passed to the application layer
type DataContextCarrier interface {
	SetRepositories(ur UserRepository, hr HealthRepository, gr GeographyRepository, tr TripRepository, cr ConversationRepository, rr RequestRepository)
	GetUserRepository() UserRepository
	GetHealthRepository() HealthRepository
	GetGeographyRepository() GeographyRepository
	GetTripRepository() TripRepository
	GetConversationRepository() ConversationRepository
	GetRequestRepository() RequestRepository
}

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	userRepository         UserRepository
	healthRepository       HealthRepository
	geographyRepository    GeographyRepository
	tripRepository         TripRepository
	conversationRepository ConversationRepository
	requestRepository      RequestRepository
}

// SetRepositories sets the repositories of the datacontext
func (dc *DataContext) SetRepositories(ur UserRepository, hr HealthRepository, gr GeographyRepository, tr TripRepository, cr ConversationRepository, rr RequestRepository) {
	dc.userRepository = ur
	dc.healthRepository = hr
	dc.geographyRepository = gr
	dc.tripRepository = tr
	dc.conversationRepository = cr
	dc.requestRepository = rr
}

// GetUserRepository returns the user repository
func (dc *DataContext) GetUserRepository() UserRepository {
	return dc.GetUserRepository()
}

// GetHealthRepository returns the health repository
func (dc *DataContext) GetHealthRepository() HealthRepository {
	return dc.healthRepository
}

// GetGeographyRepository returns the geography repository
func (dc *DataContext) GetGeographyRepository() GeographyRepository {
	return dc.GetGeographyRepository()
}

// GetTripRepository returns the trip repository
func (dc *DataContext) GetTripRepository() TripRepository {
	return dc.GetTripRepository()
}

// GetConversationRepository returns the conversation repository
func (dc *DataContext) GetConversationRepository() ConversationRepository {
	return dc.GetConversationRepository()
}

// GetRequestRepository returns the request repository
func (dc *DataContext) GetRequestRepository() RequestRepository {
	return dc.GetRequestRepository()
}
