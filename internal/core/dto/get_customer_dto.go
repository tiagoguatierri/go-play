package dto

type GetCustomerDTO struct {
	Field      string   `json:"field"`
	Value      string   `json:"value"`
	Projection []string `json:"projection"`
}
