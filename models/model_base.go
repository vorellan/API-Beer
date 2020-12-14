package models

import (
	"api-beer/utils/error_utils"
	"strings"

)

type Beer struct {
	Id       	 int64     `json:"id"`
	Name     	 string    `json:"name"`
	Brewery      string    `json:"brewery"`
	Country      string    `json:"country"`

}

func (m *Beer) Validate() error_utils.MessageErr {
	m.Name = strings.TrimSpace(m.Name)
	m.Brewery = strings.TrimSpace(m.Brewery)
	m.Country = strings.TrimSpace(m.Country)
	if m.Name == "" {
		return error_utils.NewUnprocessibleEntityError("Please enter a valid name")
	}
	if m.Brewery == "" {
		return error_utils.NewUnprocessibleEntityError("Please enter a valid brewery")
	}
	if m.Country == "" {
		return error_utils.NewUnprocessibleEntityError("Please enter a valid country")
	}
	return nil
}
