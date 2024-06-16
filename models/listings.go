package models

import "gorm.io/gorm"

type Listing struct {
	// ID uint			`gorm:"primaryKey"`
	gorm.Model
	Title        string  `json:"title" gorm:"text"`
	Description  string  `json:"description" gorm:"text"`
	Rent         float32 `json:"rent_price" gorm:"float"`
	BuildingType string  `json:"building_type" gorm:"text"`
	Availability bool    `json:"availability" gorm:"bool"`
	Location     Location
	Utility      Utility
	Features     Features
}

type Location struct {
	gorm.Model
	City         string `json:"city" gorm:"text"`
	Street       string `json:"street" gorm:"text"`
	Municipality string `json:"municipality" gorm:"text"`
	UnitNo       string `json:"unit_no" gorm:"text"`
	PostalCode   string `json:"postal_code" gorm:"text"`
	ListingID    uint   `gorm:"foreignKey:ListingID;references:Listing.ID"`
}

// type Landlord struct {
// 	gorm.Model
// 	FirstName string `json:"first_name" gorm:"text"`
// 	LastName string `json:"last_name" gorm:"text"`
// 	Email string `json:"email" gorm:"text"`
// 	Phone string `json:"phone" gorm:"text"`
// }

type Utility struct {
	ListingID uint `gorm:"foreignKey:ListingID;references:Listing.ID"`
	Water     bool `json:"water" gorm:"bool"`
	Gas       bool `json:"gas" gorm:"bool"`
	Electric  bool `json:"electric" gorm:"bool"`
	Parking   int  `json:"parking" gorm:"int"`
	Locker    bool `json:"locker" gorm:"bool"`
	Internet  bool `json:"internet" gorm:"bool"`
	Furnished bool `json:"furnished" gorm:"bool"`
}

type Features struct {
	ListingID  uint    `gorm:"foreignKey:ListingID;references:Listing.ID"`
	Rooms      int     `json:"rooms" gorm:"int"`
	Washrooms  int     `json:"washrooms" gorm:"int"`
	Kitchen    bool    `json:"kitchen" gorm:"bool"`
	SquareFeet float32 `json:"square_feet" gorm:"float"`
}
