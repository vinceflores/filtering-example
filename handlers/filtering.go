package handlers

import (
	"filtering-sample/database"
	"filtering-sample/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FilteringRequest(c *fiber.Ctx) error {

	db := database.DB.DB
	var Listing []models.Listing

	// var location models.Location
	// var utility models.Utility
	// var features models.Features

	// Get the query parameters from body
	body := new(models.FilteringRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// db.Session(&gorm.Session{})

	db.Scopes(listingFilter(body.BuildingType, body.Price)).Preload("Location", locationFilter(&body.Location)).Find(&Listing)
	

	// db.Preload(clause.Associations).Limit(10).Find(&Listing)

	return c.Status(200).JSON(Listing)
}

func listingFilter(buildingType string, price float32) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Model(&models.Listing{}).Where("Upper(building_type)like Upper(?)", buildingType).Where("rent <= ?", price)
	}
}

func locationFilter(location *models.Location) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("city = ? or city = ''", location.City).Where("postal_code = ?", location.PostalCode).Or("postal_code = ?", "").Where("street = ?", location.Street).Or("street = ?", "").Where("municipality = ?", location.Municipality).Or("municipality = ?", "")
	}
}
