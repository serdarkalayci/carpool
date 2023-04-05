// Package domain is the package that holds the very basic domain objects
package domain

// Country is the domain object for a country
type Country struct {
	ID           string
	Name         string
	Cities       []City
	BallotCities []City
}

// City is the domain object for a city
type City struct {
	Name   string
	Ballot bool
}
