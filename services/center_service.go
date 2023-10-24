package services

import (
	"errors"
	"vaccination_booking/models"
)

// id : {all details}
var Centers = make(map[string]models.Center)

// <state_name> <district_name> <center_id>
func AddVaccinationCenter(centerID string, state string, district string) {
	center := models.Center{
		CenterID:            centerID,
		State:               state,
		District:            district,
		Capacity:            make(map[int]int),
		BookingAppointments: make(map[int]map[string]bool),
	}
	Centers[centerID] = center
}

// <center_id> <day> <capacity>
func AddCapacity(centerId string, day int, capacity int) error {
	center, ok := Centers[centerId]
	if !ok {
		return errors.New("Invalid center Id") // or center not found
	}
	center.Capacity[day] = capacity
	center.BookingAppointments[day] = make(map[string]bool)
	Centers[centerId] = center
	return nil
}
