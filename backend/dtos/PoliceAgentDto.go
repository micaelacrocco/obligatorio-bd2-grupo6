package dtos

type PoliceAgentDto struct {
	CitizenID       int `json:"citizen_id"`
	PoliceStationID int `json:"police_station_id"`
	PollingPlaceID  int `json:"polling_place_id"`
}
