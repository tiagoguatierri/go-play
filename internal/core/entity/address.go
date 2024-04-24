package entity

import "errors"

var (
	ErrAddressZipCodeRequired      = errors.New("zip code is required")
	ErrAddressStreetRequired       = errors.New("street is required")
	ErrAddressNumberRequired       = errors.New("number is required")
	ErrAddressNeighborhoodRequired = errors.New("neighborhood is required")
	ErrAddressCityRequired         = errors.New("city is required")
	ErrAddressStateRequired        = errors.New("state is required")
)

type Address struct {
	ZipCode      string
	Street       string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        string
}

func NewAddress(zipCode, street, number, complement, neighborhood, city, state string) *Address {
	return &Address{
		ZipCode:      zipCode,
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
	}
}

func (a *Address) Validate() error {
	if a.ZipCode == "" {
		return ErrAddressZipCodeRequired
	}
	if a.Street == "" {
		return ErrAddressStreetRequired
	}
	if a.Number == "" {
		return ErrAddressNumberRequired
	}
	if a.Neighborhood == "" {
		return ErrAddressNeighborhoodRequired
	}
	if a.City == "" {
		return ErrAddressCityRequired
	}
	if a.State == "" {
		return ErrAddressStateRequired
	}
	return nil
}
