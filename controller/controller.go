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

// // CreateUser adds a new Account to the database
// func CreateUser(c *fiber.Ctx) error {
// 	NewUser := new(model.Account)
// 	if err := c.BodyParser(NewUser); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
// 	}
// 	if err := middleware.DBConn.Table("accounts").Create(&NewUser).Error; err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": "Failed to create NewUser"})
// 	}
// 	return c.JSON(NewUser)
// }

// func CreateUser(c *fiber.Ctx) error {
// 	NewUser := new(model.Account)
// 	if err := c.BodyParser(NewUser); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
// 	}

// 	// Create the user in the database
// 	if err := middleware.DBConn.Table("accounts").Create(&NewUser).Error; err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": "Failed to create NewUser"})
// 	}

// 	// Generate QR code for the new user (could be their email, ID, or any identifier)
// 	qrContent := fmt.Sprintf("User ID: %d, Email: %s", NewUser.ID, NewUser.EmailAddress) // Customize as needed
// 	qrCodePath := fmt.Sprintf("./qrcodes/user_%d.png", NewUser.ID)

// 	// Generate and save the QR code
// 	if err := qrcode.WriteFile(qrContent, qrcode.Medium, 256, qrCodePath); err != nil {
// 		log.Println("Failed to generate QR code:", err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate QR code"})
// 	}

// 	// Optionally, attach the QR code URL to the response (if you're serving it from a web server)
// 	qrCodeURL := fmt.Sprintf("http://localhost:3000/qrcodes/user_%d.png", NewUser.ID)

// 	// Return the new user with the QR code URL
// 	return c.JSON(fiber.Map{
// 		"user":      NewUser,
// 		"qrCodeURL": qrCodeURL,
// 	})
// }

// //Login

// // CreateApprovalStatus creates an approval status record for a specific user
// func CreateApprovalStatus(c *fiber.Ctx) error {
// 	var input model.ApprovalStatus

// 	// Bind JSON to struct
// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	// Check if account exists
// 	var account model.Account
// 	if err := middleware.DBConn.First(&account, "id = ?", input.AccountID).Error; err != nil {
// 		return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 			"error": "Account not found",
// 		})
// 	}

// 	// Prevent duplicate approval status
// 	var existing model.ApprovalStatus
// 	if err := middleware.DBConn.Where("account_id = ?", input.AccountID).First(&existing).Error; err == nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Approval status already exists for this user",
// 		})
// 	}

// 	// Default status
// 	if input.Status == "" {
// 		input.Status = "Pending"
// 	}

// 	// Create record
// 	if err := middleware.DBConn.Create(&input).Error; err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to create approval status",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Approval status created",
// 		"data":    input,
// 	})
// }

// // ApproveUser updates approval status to "Approved" if all 4 requirements are true
// func ApproveUser(c *fiber.Ctx) error {
// 	accountIDParam := c.Params("id")

// 	// Parse UUID
// 	accountID, err := uuid.Parse(accountIDParam)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid account ID format",
// 		})
// 	}

// 	// Find the approval status by account ID
// 	var status model.ApprovalStatus
// 	if err := middleware.DBConn.Where("account_id = ?", accountID).First(&status).Error; err != nil {
// 		return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 			"error": "Approval status not found for this user",
// 		})
// 	}

// 	// Check all 4 requirements
// 	if status.Requirement1 && status.Requirement2 && status.Requirement3 && status.Requirement4 {
// 		status.Status = "Approved"
// 		status.IsApproved = true

// 		// Save changes
// 		if err := middleware.DBConn.Save(&status).Error; err != nil {
// 			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to approve user",
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"message": "User approved successfully",
// 			"status":  status,
// 		})
// 	}

// 	// If not all requirements are met
// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 		"error": "User cannot be approved. Not all requirements are met.",
// 		"status": fiber.Map{
// 			"requirement_1": status.Requirement1,
// 			"requirement_2": status.Requirement2,
// 			"requirement_3": status.Requirement3,
// 			"requirement_4": status.Requirement4,
// 		},
// 	})
// }

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