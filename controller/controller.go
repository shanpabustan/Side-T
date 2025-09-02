package controller

import (
	//"fmt"
	"context"
	"intern_template_v1/middleware"
	"intern_template_v1/model"
	"strconv"
	"time"

	//"log"
	//"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	//"github.com/skip2/go-qrcode"
	//"github.com/google/uuid"
)



func AddItem(c *fiber.Ctx) error {
    // Parse regular fields (not the file)
    newItem := new(model.Items)
    newItem.ProductName = c.FormValue("product_name")
	newItem.Category = c.FormValue("category")

	newItem.ProductName = c.FormValue("product_name")
    newItem.Category = c.FormValue("category")

    price, _ := strconv.ParseFloat(c.FormValue("price"), 64)
    newItem.Price = price

    qty, _ := strconv.Atoi(c.FormValue("quantity"))
    newItem.Quantity = qty
    // Check if file exists
    fileHeader, err := c.FormFile("link")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Image file is required"})
    }

    // Open file
    file, err := fileHeader.Open()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to open image"})
    }
    defer file.Close()

    cld, err := cloudinary.New()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to initialize Cloudinary"})
		}

    // Upload to Cloudinary
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
        Folder: "items",
    })
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to upload to Cloudinary"})
    }

    // Save only the Cloudinary URL into Postgres
    newItem.Link = uploadResult.SecureURL

    // Save item data to DB
    if err := middleware.DBConn.Table("items").Create(&newItem).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create item"})
    }

    return c.JSON(newItem)
}


func GetItems(c *fiber.Ctx) error {
	var items []model.Items

	// Fetch all items from DB
	if err := middleware.DBConn.Table("items").Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch items"})
	}

	// Return as JSON
	return c.JSON(items)
}


func AddPurchase(c *fiber.Ctx) error {
	// Parse input (item name & quantity)
	var input struct {
		ItemName string `json:"item_name"`
		Quantity int    `json:"quantity"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Find item in DB
	var item model.Items
	if err := middleware.DBConn.Where("product_name = ?", input.ItemName).First(&item).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	// Check stock
	if item.Quantity < input.Quantity {
		return c.Status(400).JSON(fiber.Map{"error": "Not enough stock"})
	}

	// Reduce item quantity
	item.Quantity -= input.Quantity
	if err := middleware.DBConn.Save(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update item stock"})
	}

	// Create purchase record
	purchase := model.Purchases{
		ItemName:   item.ProductName,
		Quantity:   input.Quantity,
		TotalPrice: float64(input.Quantity) * item.Price,
	}

	if err := middleware.DBConn.Create(&purchase).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to record purchase"})
	}

	return c.JSON(purchase)
}
