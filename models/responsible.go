package models

type Responsible struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CPF           string `json:"cpf"`
	Street        string `json:"street"`
	Number        string `json:"number"`
	ZIP           string `json:"zip"`
	Complement    string `json:"complement"`
	Status        string `json:"status" validate:"oneof=OK ACTIVE BLOCKED BANNED"`
	CardToken     string `json:"card_token"`
	PaymentMethod string `json:"payment_method"`
	CustomerId    string `json:"customer_id"`
}
