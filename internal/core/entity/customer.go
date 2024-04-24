package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNameRequired  = errors.New("name is required")
	ErrCustomerEmailRequired = errors.New("email is required")
	ErrCustomerCpfRequired   = errors.New("cpf is required")
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	Cpf       string
	Age       uint
	Addresses []*Address
}

type CustomerOption func(*Customer)

func NewCustomer(name, email, cpf string, age uint, opts ...CustomerOption) *Customer {
	c := &Customer{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
		Cpf:   cpf,
		Age:   age,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithAddresses(addresses ...*Address) CustomerOption {
	return func(c *Customer) {
		c.Addresses = append(c.Addresses, addresses...)
	}
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return ErrCustomerNameRequired
	}

	if c.Email == "" {
		return ErrCustomerEmailRequired
	}

	if c.Cpf == "" {
		return ErrCustomerCpfRequired
	}

	for _, a := range c.Addresses {
		if err := a.Validate(); err != nil {
			return err
		}
	}

	return nil
}
