package services

import (
	"errors"
	"vaccination_booking/models"
)

// <center_id> <day> <user_id>
func BookVaccination(centerID string, day int, userID string) (models.BookVaccinationResponse, error) { // response or true/false
	res := models.BookVaccinationResponse{
		Message: ""
	}
	center, ok := Centers[centerID]
	if !ok {
		return res, errors.New("Invalid center Id") // or center not found
	}
	
	if _, booked := center.BookingAppointments[day][userID]; booked {
		return res, errors.New("Already booked")
	}
	
	// sync
	if len(center.BookingAppointments[day]) >= center.Capacity[day] {
		return res, errors.New("No free slots")
	}

	center.BookingAppointments[day][userID] = true
	Centers[centerID] = center

	res.Message = "Successfully Booked"
	return res, nil
}
