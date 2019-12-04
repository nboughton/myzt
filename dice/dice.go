package dice

import (
	"strconv"

	"github.com/nboughton/go-roll"
)

// D66 rolls a D66
func D66() int {
	n, _ := strconv.Atoi(roll.D6.Roll().Value + roll.D6.Roll().Value)
	return n
}

// D666 rolls a D666
func D666() int {
	n, _ := strconv.Atoi(roll.D6.Roll().Value + roll.D6.Roll().Value + roll.D6.Roll().Value)
	return n
}
