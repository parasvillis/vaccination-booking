package services

import (
	"errors"
	"vaccination_booking/models"
)

/*
	"id" : {
		all details
	}
*/

var Users = make(map[string]models.User)

// <unique_identification> <name> <gender> <age> <current_state> <current_district>

func AddUser(id string, name string, gender string, age int, currentState string, currentDistrict string) error {

	if age < 18 {
		return errors.New("Below 18 years., not eligible for vaccination")
	}

	user := models.User{
		ID:              id,
		Name:            name,
		Gender:          gender,
		Age:             age,
		CurrentState:    currentState,
		CurrentDistrict: currentDistrict,
	}
	Users[id] = user
	return nil
}
