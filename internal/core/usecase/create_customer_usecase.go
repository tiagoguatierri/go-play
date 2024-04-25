package usecase

import (
	"fmt"

	"github.com/tiagoguatierri/go-play/internal/core/entity"
	"github.com/tiagoguatierri/go-play/internal/core/repository"
)

type CreateCustomerRequest struct {
	Name  string
	Cpf   string
	Email string
	Age   uint
}

type CreateCustomerUseCase struct {
	repo repository.Repository[*entity.Customer]
}

func NewCreateCustomerUseCase(repo repository.Repository[*entity.Customer]) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{repo}
}

func (us *CreateCustomerUseCase) Execute(req CreateCustomerRequest) (*entity.Customer, error) {
	customer := entity.NewCustomer(req.Name, req.Email, req.Cpf, req.Age)

	if err := customer.Validate(); err != nil {
		return nil, fmt.Errorf("CreateCustomerUseCaseError: %w", err)
	}

	err := us.repo.Create(customer)
	if err != nil {
		return nil, fmt.Errorf("CreateCustomerUseCaseError: %w", err)
	}

	return customer, nil
}
