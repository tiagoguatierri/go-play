package usecase

import (
	"fmt"

	"github.com/tiagoguatierri/go-play/internal/core/entity"
	"github.com/tiagoguatierri/go-play/internal/core/repository"
)

type GetCustomerRequest struct {
	Field      entity.CustomerField
	Value      string
	Projection []string
}

type GetCustomerUseCase struct {
	repo repository.Repository[*entity.Customer]
}

func NewGetCustomerUseCase(repo repository.Repository[*entity.Customer]) *GetCustomerUseCase {
	return &GetCustomerUseCase{repo}
}

func (us *GetCustomerUseCase) Execute(req GetCustomerRequest) (*entity.Customer, error) {
	c := new(entity.Customer)

	if req.Field == entity.CustomerFieldID {
		c.ID = req.Value
	}

	if req.Field == entity.CustomerFieldEmail {
		c.Email = req.Value
	}

	if req.Field == entity.CustomerFieldCpf {
		c.Cpf = req.Value
	}

	customer, err := us.repo.Get(c)
	if err != nil {
		return nil, fmt.Errorf("GetCustomerUseCaseError: %w", err)
	}

	return customer, nil
}
