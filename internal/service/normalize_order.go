package service

import (
	"strconv"
	"strings"
	"technical-skills/internal/model"
)

// splitProductID extracts product items from a string separated by "/"
// and removes unnecessary special characters such as "%20x", "%20", and leading "-"
// and removes that is not related to the product name and quantity ordered.
//
// Parameters:
//   - inputProduct: A string containing product items separated by "/".
//     Example: " --FG0A-CLEAR-OPPOA3*2/%20xFG0A-MATTE-OPPOA3 "
//     Output: ["FG0A-CLEAR-OPPOA3*2", "FG0A-MATTE-OPPOA3"]
//
// Returns:
//   - []string: A slice containing individual product names.
func splitProductID(inputProduct string) []string {
	// I avoid using regex in this step because, although it is flexible,
	// it is generally slower compared to simple string operations.
	replacer := strings.NewReplacer("%20x", "", "%20", "")
	inputProduct = replacer.Replace(inputProduct)

	// Split input product string by "/"
	result := strings.Split(inputProduct, "/")

	// Preallocate the slice to improve performance and avoid unnecessary memory allocations
	products := make([]string, 0, len(result))

	// Remove unnecessary characters from product names and append to products slice
	for _, r := range result {
		r = strings.TrimSpace(r)      // Remove leading and trailing spaces
		r = strings.TrimLeft(r, "-")  // Remove leading
		r = strings.TrimRight(r, "-") // Remove trailing

		// Remove everything before "FG0A"
		if index := strings.Index(r, "FG0A"); index != -1 {
			r = r[index:]
		}

		products = append(products, r) // Store cleaned product name
	}

	return products
}

// extractQuantity extracts the product ID and quantity from a given string.
//
// Parameters:
//   - product: A string in the format "{filmTypeID}-{textureID}-{phoneModelID}*n"
//     Example 1: "FG0A-MATTE-IPHONE16PROMAX*3" → productID: "FG0A-MATTE-IPHONE16PROMAX", quantity: 3
//     Example 1: "3*FG0A-MATTE-IPHONE16PROMAX" → productID: "FG0A-MATTE-IPHONE16PROMAX", quantity: 3
//
// Returns:
//   - productID: The extracted product ID.
//   - quantity: The specified quantity (defaults to 1 if not provided).
func extractQuantity(product string) (productID string, quantity int) {
	parts := strings.Split(product, "*")

	if len(parts) == 2 {
		// Check if the second part is a number
		if qty, err := strconv.Atoi(parts[1]); err == nil {
			return parts[0], qty
		}

		// Check if the first part is a number
		if qty, err := strconv.Atoi(parts[0]); err == nil {
			return parts[1], qty
		}
	}

	// If no "*" is present or no valid number is found, return the product as is with quantity 1
	return product, 1
}

// splitProductDetail splits a product string into three parts:
// 1. Film Type ID (e.g., FG0A, FG05)
// 2. Texture ID (e.g., CLEAR, MATTE, PRIVACY)
// 3. Phone Model ID (e.g., IPHONE16PROMAX, SAMSUNGS25, OPPOA3)
//
// Parameters:
//   - product: A string containing the product code in the format "{filmTypeID}-{textureID}-{phoneModelID}"
//     Example: "FG0A-PRIVACY-IPHONE16PROMAX-B" → filmTypeID: "FG0A", textureID: "PRIVACY", phoneModelID: "IPHONE16PROMAX-B"
//
// Returns:
//   - filmTypeID: The film type ID.
//   - textureID: The texture ID.
//   - phoneModelID: The phone model ID.
func splitProductDetail(product string) (filmTypeID, textureID, phoneModelID string) {
	parts := strings.SplitN(product, "-", 3)
	if len(parts) > 0 {
		filmTypeID = parts[0]
	}
	if len(parts) > 1 {
		textureID = parts[1]
	}
	if len(parts) > 2 {
		phoneModelID = parts[2]
	}
	return
}

func NormalizeOrder(inputOrders []model.InputOrder) (cleanedOrders []model.CleanedOrder) {
	// Stores the names of complimentary cleanners and their quantities
	cleannerMap := make(map[string]int)

	// Stores the order of complimentary cleanners
	cleannerMapOrder := []string{}

	// Total quantity of products in the order
	totalProductQuantity := 0

	// Stores the No.
	no := 1
	for _, inputOrder := range inputOrders {
		// Extract product IDs from the platform product ID
		products := splitProductID(inputOrder.PlatformProductId)

		// Calculate total ordered quantity
		total := 0
		productQuantities := make([]int, len(products))
		productIDs := make([]string, len(products))
		for i, product := range products {
			productID, quantity := extractQuantity(product)
			productIDs[i] = productID
			productQuantities[i] = inputOrder.Qty * quantity
			total += productQuantities[i]
		}
		totalProductQuantity += total

		// Calculate unit price by dividing total price by total quantity (avoid division by zero)
		unitPrice := 0
		if total > 0 {
			unitPrice = inputOrder.TotalPrice / total
		}

		for i := range products {
			// productID, _ := extractQuantity(product)
			quantity := productQuantities[i]
			totalPrice := unitPrice * quantity

			// Find the film type, texture, and phone model id to normalize form
			filmTypeID, textureID, phoneModelID := splitProductDetail(productIDs[i])
			cleanedOrders = append(cleanedOrders, model.CleanedOrder{
				No:         no,
				ProductId:  filmTypeID + "-" + textureID + "-" + phoneModelID,
				MaterialId: filmTypeID + "-" + textureID,
				ModelId:    phoneModelID,
				Qty:        quantity,
				UnitPrice:  unitPrice,
				TotalPrice: totalPrice,
			})
			no++

			// Store the type of cleaner and the quantity to be given as a freebie
			cleannerKey := textureID + "-CLEANNER"
			if qty, exists := cleannerMap[cleannerKey]; exists {
				cleannerMap[cleannerKey] = qty + quantity
			} else {
				cleannerMap[cleannerKey] = quantity
				cleannerMapOrder = append(cleannerMapOrder, cleannerKey)
			}
		}
	}

	// Append wiping cloth equal to total ordered quantity
	cleanedOrders = append(cleanedOrders, model.CleanedOrder{
		No:        no,
		ProductId: "WIPING-CLOTH",
		Qty:       totalProductQuantity,
	})
	no++

	// Append complimentary cleaners in recorded order
	for _, key := range cleannerMapOrder {
		cleanedOrders = append(cleanedOrders, model.CleanedOrder{
			No:        no,
			ProductId: key,
			Qty:       cleannerMap[key],
		})
		no++
	}

	return
}
