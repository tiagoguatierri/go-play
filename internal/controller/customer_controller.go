package controller

import (
	"fmt"

	"github.com/tiagoguatierri/go-play/internal/core/dto"
	"github.com/tiagoguatierri/go-play/internal/core/entity"
	"github.com/tiagoguatierri/go-play/internal/core/repository"
	"github.com/tiagoguatierri/go-play/internal/core/usecase"
)

type CustomerController struct {
	createCustomerUseCase *usecase.CreateCustomerUseCase
	getCustomerUseCase    *usecase.GetCustomerUseCase
}

func NewCustomerController(repo repository.Repository[*entity.Customer]) *CustomerController {
	return &CustomerController{
		createCustomerUseCase: usecase.NewCreateCustomerUseCase(repo),
		getCustomerUseCase:    usecase.NewGetCustomerUseCase(repo),
	}
}

func (ctrl *CustomerController) Create(dto dto.CreateCustomerDTO) (*entity.Customer, error) {
	gcr := usecase.GetCustomerRequest{}
	gcr.Field = entity.CustomerFieldCpf
	gcr.Value = dto.Cpf
	exists, err := ctrl.getCustomerUseCase.Execute(gcr)

	if err != nil {
		return nil, fmt.Errorf("CustomerControllerError: %w", err)
	}

	if exists != nil {
		return nil, entity.ErrCustomerAlreadyExists
	}

	ccr := usecase.CreateCustomerRequest{}
	ccr.Name = dto.Name
	ccr.Cpf = dto.Cpf
	ccr.Email = dto.Email
	ccr.Age = dto.Age

	customer, err := ctrl.createCustomerUseCase.Execute(ccr)
	if err != nil {
		return nil, fmt.Errorf("CustomerControllerError: %w", err)
	}

	return customer, nil
}

func (ctrl *CustomerController) Get(dto dto.GetCustomerDTO) (*entity.Customer, error) {
	req := usecase.GetCustomerRequest{}
	req.Field = entity.CustomerField(dto.Field)
	req.Value = dto.Value
	req.Projection = dto.Projection

	customer, err := ctrl.getCustomerUseCase.Execute(req)
	if err != nil {
		return nil, fmt.Errorf("CustomerControllerError: %w", err)
	}

	return customer, nil
}
