package handlers

import (
	"github.com/pandeyvijay99/gohummbird-teacher/db"
	"github.com/pandeyvijay99/gohummbird-teacher/models"

	"github.com/gofiber/fiber"
)

// GetAllProducts query all products
func GetAllProducts(c *fiber.Ctx) {
	db := db.DB
	var products []models.Product
	db.Find(&products)
	c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// GetProduct query product
func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DB
	var product models.Product
	db.Find(&product, id)
	if product.Title == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

// CreateProduct new product
func CreateProduct(c *fiber.Ctx) {
	db := db.DB
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
		return
	}
	db.Create(&product)
	c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DB

	var product models.Product
	db.First(&product, id)
	if product.Title == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
		return
	}
	db.Delete(&product)
	c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
