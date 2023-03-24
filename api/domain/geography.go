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
