package dto

type CreateCustomerDTO struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}
