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
	getReq := usecase.GetCustomerRequest{}
	getReq.Field = entity.CustomerFieldCpf
	getReq.Value = dto.Cpf
	exists, err := ctrl.getCustomerUseCase.Execute(getReq)

	if err != nil {
		return nil, fmt.Errorf("CustomerControllerError: %w", err)
	}

	if exists != nil {
		return nil, entity.ErrCustomerAlreadyExists
	}

	createReq := usecase.CreateCustomerRequest{}
	createReq.Name = dto.Name
	createReq.Cpf = dto.Cpf
	createReq.Email = dto.Email
	createReq.Age = dto.Age

	customer, err := ctrl.createCustomerUseCase.Execute(createReq)
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
