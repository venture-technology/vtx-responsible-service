package models

type Child struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	RG          string      `json:"rg"`
	Responsible Responsible `json:"responsible"`
}
