package main

import (
	"fmt"
	"strings"
)

var productPrice = map[string]float64{
	"TShirt": 20.00,
	"MUG":    5.00,
	"HAT":   15.00,
	"BOOK":  10.00,
}

func calulateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrice[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrice[originalItemCode]
			if found {
				salePrice := basePrice * 0.8 // 20% discount
				fmt.Printf("Item: %s, Original Price: %.2f, Sale Price: %.2f\n", originalItemCode, basePrice, salePrice)
			}
			return basePrice * 0.8, found
		}
		fmt.Printf("Item code %s not found.\n", itemCode)
		return 0.0, false
	}
	return basePrice, true
}
func main() {

	fmt.Println("---------------- Simple Sales System ----------------")
	orderItems := []string{"TShirt", "MUG_SALE", "HAT", "BOOK_SALE", "UNKNOWN_ITEM"}
	for i := 0; i < len(orderItems); i++ {
		itemCode := orderItems[i]
		price, found := calulateItemPrice(itemCode)
		fmt.Printf("Item: %s, Price: %.2f, Found: %t\n", itemCode, price, found)
	}
}