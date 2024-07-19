package models

type Responsible struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CPF        string `json:"cpf"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	ZIP        string `json:"zip"`
	Complement string `json:"complement"`
}
