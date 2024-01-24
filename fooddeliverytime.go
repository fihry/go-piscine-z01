package piscine

// burger takes 15 min, chips takes 10 min and nuggets takes 12 min

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		return 15
	}
	if order == "chips" {
		return 10
	}
	if order == "nuggets" {
		return 12
	} else {
		return 404
	}
}
