package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tiagoguatierri/go-play/internal/core/entity"
)

type AddressTestSuite struct {
	suite.Suite
	Address *entity.Address
}

// like a "beforeEach"
func (suite *AddressTestSuite) SetupTest() {
	suite.Address = entity.NewAddress("12345678", "Rua dos Bobos", "0", "Apto 101", "Centro", "São Paulo", "SP")
}

func (suite *AddressTestSuite) TestAddressValid() {
	err := suite.Address.Validate()
	assert.Nil(suite.T(), err)
}

func (suite *AddressTestSuite) TestAddressInvalid() {
	suite.Address.ZipCode = ""
	err := suite.Address.Validate()

	assert.Equal(suite.T(), entity.ErrAddressZipCodeRequired, err)
}

func TestAddressTestSuite(t *testing.T) {
	suite.Run(t, new(AddressTestSuite))
}
