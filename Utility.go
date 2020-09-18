package OctaForceEngine

import (
	"strconv"
)

func parseFloat(number string) float32 {
	floatVar, _ := strconv.ParseFloat(number, 32)
	return float32(floatVar)
}
func parseInt(number string) uint32 {
	intVar, _ := strconv.ParseInt(number, 10, 32)
	return uint32(intVar)
}
