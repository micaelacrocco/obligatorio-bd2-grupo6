package dtos

type PollingPlaceDto struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	ZoneID  int    `json:"zone_id"`
}
