package domain

type Country struct {
	ID           string
	Name         string
	Cities       []City
	BallotCities []City
}

type City struct {
	Name   string
	Ballot bool
}

type ErrInvalidDestination struct{}

func (e ErrInvalidDestination) Error() string {
	return "destination is not a ballot city"
}
