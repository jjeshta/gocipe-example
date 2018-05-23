// generated by gocipe 2684eeb0295ee5f2f4177b7c406ae5c8553e74d3b44b524a3eeb29e37d1363f0; DO NOT EDIT

package citizen

import (
	"time"

	"github.com/fluxynet/gocipe-example/models/country"
)

// Citizen A human being belonging to a country
type Citizen struct {
	ID         *string          `json:"id"`
	Surname    *string          `json:"surname"`
	OtherNames *string          `json:"othernames"`
	Gender     *string          `json:"gender"`
	DOB        *time.Time       `json:"dob"`
	Country    *country.Country `json:"country"`
}

// New returns an instance of Citizen
func New() *Citizen {
	return &Citizen{
		ID:         new(string),
		Surname:    new(string),
		OtherNames: new(string),
		Gender:     new(string),
		DOB:        new(time.Time),
	}
}
