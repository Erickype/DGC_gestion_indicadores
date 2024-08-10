package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterPeopleRequest struct {
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
}

func FilterPeople(people *[]Person, filterPeopleRequest *FilterPeopleRequest) (err error) {
	query := database.DB

	// Lista de condiciones OR
	var conditions []string
	var values []interface{}

	// Agrega las condiciones de búsqueda "LIKE" para cada campo que no esté vacío
	if filterPeopleRequest.Identity != "" {
		conditions = append(conditions, "LOWER(identity) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Identity)))
	}
	if filterPeopleRequest.Name != "" {
		conditions = append(conditions, "LOWER(name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Name)))
	}
	if filterPeopleRequest.Lastname != "" {
		conditions = append(conditions, "LOWER(lastname) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Lastname)))
	}
	if filterPeopleRequest.Email != "" {
		conditions = append(conditions, "LOWER(email) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Email)))
	}

	// Si hay condiciones, se unen con OR
	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	// Ejecuta la consulta final
	err = query.Limit(model.FilterSize).Find(people).Error
	if err != nil {
		return err
	}
	return nil
}
