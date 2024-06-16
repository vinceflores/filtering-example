package models

type FilteringRequest struct {
	Location     Location `json:"location"`
	Utility      Utility  `json:"utility"`
	Features     Features `json:"features"`
	Price        float32  `json:"price"`
	BuildingType string   `json:"building_type"`
}
