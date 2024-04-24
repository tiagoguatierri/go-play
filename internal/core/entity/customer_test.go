package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tiagoguatierri/go-play/internal/core/entity"
)

type CustomerTestSuite struct {
	suite.Suite
	Customer              *entity.Customer
	CustomerWithAddresses *entity.Customer
}

// like a "beforeEach"
func (suite *CustomerTestSuite) SetupTest() {
	suite.Customer = entity.NewCustomer("Tiago", "tiagovit@gmail.com", "12345678901", 33)
	suite.CustomerWithAddresses = entity.NewCustomer(
		"Tiago",
		"tiagovit@gmail.com",
		"12345678901",
		33,
		entity.WithAddresses(
			entity.NewAddress("12345678", "Rua dos Bobos", "0", "Apto 101", "Centro", "São Paulo", "SP"),
		),
	)
}

func (suite *CustomerTestSuite) TestCustomerValid() {
	err := suite.Customer.Validate()
	assert.Nil(suite.T(), err)
}

func TestCustomerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerTestSuite))
}
