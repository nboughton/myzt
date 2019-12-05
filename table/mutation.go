package table

import "github.com/nboughton/go-roll"

// Mutation roll table
var Mutation = roll.List{
	Name: "Mutations",
	Items: []string{
		"Acid Spit",
		"Amphibian",
		"Corpse-Eater",
		"Extreme Reflexes",
		"Flame Breather",
		"Four Armed",
		"Frog Legs",
		"Human Magnet",
		"Human Plant",
		"Insectoid",
		"Insect Wings",
		"Luminescence",
		"Manbeast",
		"Mind Terror",
		"Puppeteer",
		"Parasite",
		"Pathokinesis",
		"Pyrokinesis",
		"Reptilian",
		"Rot-Eater",
		"Sonar",
		"Spores",
		"Sprinter",
		"Telepathy",
		"Tracker",
	},
}

// MutationMisfire table
var MutationMisfire = roll.List{
	Items: []string{
		"Your mutant powers run amok, and you suffer one point of permanent trauma. You choose which attribute is affected. At the same time, you develop a new mutation.",
		"You suffer the effect of the mutation, and take as much trauma as your target. If the mutation doesn’t cause trauma, you get disoriented instead and can’t act at all in the next turn.",
		"The mutation consumes twice the amount of MP you intended to spend on it – without increasing the effect. You can’t drop below zero MP.",
		"The mutation locks down after this use. You can’t use it again for the rest of the session.",
		"The mutation changes your appearance in some way. You choose how. The effect is only cosmetic, but it is permanent.",
		"The mutation gets supercharged. You get back the MP you just spent, and you can immediately – in the same turn –  activate the same mutation again, against the same target or another.",
	},
}
