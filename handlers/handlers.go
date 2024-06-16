package handlers

import (
	"filtering-sample/database"
	"filtering-sample/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)



func CreateListing(c *fiber.Ctx) error {
	db := database.DB.DB
	body := new(models.Listing)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	listings := models.Listing{
		Title:        body.Title,
		Description:  body.Description,
		Rent:         body.Rent,
		BuildingType: body.BuildingType,
		Availability: body.Availability,
	}

	db.Omit(clause.Associations).Create(&listings)

	if db.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": db.Error,
		})
	}
	location := models.Location{
		City:         body.Location.City,
		Street:       body.Location.Street,
		Municipality: body.Location.Municipality,
		UnitNo:       body.Location.UnitNo,
		PostalCode:   body.Location.PostalCode,
		ListingID:    listings.ID,
	}
	utility := models.Utility{
		Water:     body.Utility.Water,
		Gas:       body.Utility.Gas,
		Electric:  body.Utility.Electric,
		Parking:   body.Utility.Parking,
		Locker:    body.Utility.Locker,
		Internet:  body.Utility.Internet,
		Furnished: body.Utility.Furnished,
		ListingID: listings.ID,
	}
	features := models.Features{
		Rooms:      body.Features.Rooms,
		Washrooms:  body.Features.Washrooms,
		Kitchen:    body.Features.Kitchen,
		SquareFeet: body.Features.SquareFeet,
		ListingID:  listings.ID,
	}

	db.Create(&utility)
	db.Create(&features)
	db.Create(&location)

	listings.Location = location
	listings.Utility = utility
	listings.Features = features

	db.Save(&listings)
	return c.Status(200).JSON(listings)
}

func GetListingDetails(c *fiber.Ctx) error {
	// Get the ID from the URL
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}
	db := database.DB.DB
	var listing models.Listing
	err := db.Preload(clause.Associations).First(&listing, id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"listing": listing,
		"id":      id,
	})
}
