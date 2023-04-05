package application

// MockContext mimics the DataContext struct, which is used to test the application layer
type MockContext struct {
	userRepository         UserRepository
	healthRepository       HealthRepository
	geographyRepository    GeographyRepository
	tripRepository         TripRepository
	conversationRepository ConversationRepository
	requestRepository      RequestRepository
}

func (mc *MockContext) SetRepositories(ur UserRepository, hr HealthRepository, gr GeographyRepository, tr TripRepository, cr ConversationRepository, rr RequestRepository) {
	mc.userRepository = ur
	mc.healthRepository = hr
	mc.geographyRepository = gr
	mc.tripRepository = tr
	mc.conversationRepository = cr
	mc.requestRepository = rr
}

func (mc *MockContext) GetHealthRepository() HealthRepository {
	return mc.healthRepository
}

func (mc *MockContext) GetTripRepository() TripRepository {
	return mc.tripRepository
}

func (mc *MockContext) GetRequestRepository() RequestRepository {
	return mc.requestRepository
}

func (mc *MockContext) GetConversationRepository() ConversationRepository {
	return mc.conversationRepository
}

func (mc *MockContext) GetGeographyRepository() GeographyRepository {
	return mc.geographyRepository
}

func (mc *MockContext) GetUserRepository() UserRepository {
	return mc.userRepository
}
