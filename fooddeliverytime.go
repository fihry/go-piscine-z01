package piscine

// burger takes 15 min, chips takes 10 min and nuggets takes 12 min
type food struct {
	preptime int
}

var foods = map[string]food{
	"burger": {15}, "chips": {10}, "nuggets": {12},
}

func FoodDeliveryTime(order string) int {
	return foods[order].preptime
}
