package table

import "github.com/nboughton/go-roll"

// ArkEvents to roll for
var ArkEvents = roll.Table{
	Name: "Ark Events",
	Dice: roll.Dice{1, roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12, 13, 14}, Text: "Danger Ahead"},
		{Match: []int{15, 16, 21, 22}, Text: "Fight over projects"},
		{Match: []int{23, 24, 25, 26}, Text: "A threat escalates"},
		{Match: []int{31, 32, 33}, Text: "Fight over gear"},
		{Match: []int{34, 35, 36, 41}, Text: "An appealing offer"},
		{Match: []int{42, 43, 44, 45}, Text: "A PC gets into trouble"},
		{Match: []int{46, 51, 52, 53}, Text: "An NPC gets into trouble"},
		{Match: []int{54, 55, 56}, Text: "PC vs PC"},
		{Match: []int{61, 62, 63}, Text: "Unwanted consequences"},
		{Match: []int{64, 65, 66}, Text: "A glimpse of the dream"},
	},
}

// ZoneEvents to roll for
var ZoneEvents = roll.Table{
	Name: "Zone Events",
	Dice: roll.Dice{1, roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12, 13, 14}, Text: "A beacon in the distance"},
		{Match: []int{15, 16, 21, 22}, Text: "Danger ahead"},
		{Match: []int{23, 24, 25}, Text: "A hard choice"},
		{Match: []int{26, 31, 32, 33}, Text: "Gear is lost"},
		{Match: []int{34, 35, 36}, Text: "Sleepless night"},
		{Match: []int{41, 42, 43, 44}, Text: "A PC gets into trouble"},
		{Match: []int{45, 46, 51, 52}, Text: "An NPC gets into trouble"},
		{Match: []int{53, 54, 55}, Text: "PC vs PC"},
		{Match: []int{56, 61, 62}, Text: "The zone vs everyone"},
		{Match: []int{63, 64, 65, 66}, Text: "An unexpected encounter"},
	},
}
