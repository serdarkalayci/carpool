// Package application is the package that holds the application logic between database and communication layers
package application

// HealthRepository is the interface to interact with database
type HealthRepository interface {
	Ready() bool
}

// HealthService is the struct to let outer layers to interact to the health applicatopn
type HealthService struct {
	dc DataContextCarrier
}

// NewHealthService creates a new HealthService instance and sets its repository
func NewHealthService(dc DataContextCarrier) HealthService {
	return HealthService{
		dc: dc,
	}
}

// Ready returns true if underlying reposiroty and its connection is up and running, false otherwise
func (hs HealthService) Ready() bool {
	return hs.dc.GetHealthRepository().Ready()
}
