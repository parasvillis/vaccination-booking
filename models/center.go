package models

// <state_name> <district_name> <center_id>

type Center struct {
	CenterID            string
	State               string
	District            string
	Capacity            map[int]int             // store capacity for each day
	BookingAppointments map[int]map[string]bool // ref notepad please L 21 :)
}

type BookVaccinationResponse struct {
	Message string // we can have respCodes here also
}
