package table

import (
	"strconv"
	"strings"

	"github.com/nboughton/go-roll"
	"github.com/nboughton/swnt/content/format"
)

// NPC type
type NPC struct {
	Stats      npcStatBlock
	Name       string
	Appearance string
	Goal       string
	Weapon     string
}

// NewNPC roll an npcGen table
func NewNPC(t format.OutputType, npcType string) NPC {
	var n NPC

	switch strings.ToLower(npcType) {
	case "enforcer":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Enforcer.Name.Roll(),
			Appearance: Enforcer.Appearance.Roll(),
			Goal:       Enforcer.Goal.Roll(),
			Weapon:     Enforcer.Weapon.Roll(),
		}
	case "gearhead":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Gearhead.Name.Roll(),
			Appearance: Gearhead.Appearance.Roll(),
			Goal:       Gearhead.Goal.Roll(),
			Weapon:     Gearhead.Weapon.Roll(),
		}
	case "stalker":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Stalker.Name.Roll(),
			Appearance: Stalker.Appearance.Roll(),
			Goal:       Stalker.Goal.Roll(),
			Weapon:     Stalker.Weapon.Roll(),
		}
	case "fixer":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Fixer.Name.Roll(),
			Appearance: Fixer.Appearance.Roll(),
			Goal:       Fixer.Goal.Roll(),
			Weapon:     Fixer.Weapon.Roll(),
		}
	default:
	}

	return n
}

type npcStatBlock struct {
	Role     string
	Strength int
	Agility  int
	Wits     int
	Empathy  int
	Skills   string
}

var npcStatBlockHeaders = []string{"Role", "Strength", "Agility", "Wits", "Empathy", "Skills"}

func (n npcStatBlock) Format(t format.OutputType) string {
	return format.Table(t,
		npcStatBlockHeaders,
		[][]string{{n.Role, strconv.Itoa(n.Strength), strconv.Itoa(n.Agility), strconv.Itoa(n.Wits), strconv.Itoa(n.Empathy), n.Skills}},
	)
}

func (n npcStatBlock) String() string {
	return n.Format(format.TEXT)
}

type npcStatBlockTable []npcStatBlock

func (n npcStatBlockTable) Format(t format.OutputType) string {
	rows := [][]string{}

	for _, row := range n {
		rows = append(rows, []string{row.Role, strconv.Itoa(row.Strength), strconv.Itoa(row.Agility), strconv.Itoa(row.Wits), strconv.Itoa(row.Empathy), row.Skills})
	}

	return format.Table(t, npcStatBlockHeaders, rows)
}

func (n npcStatBlockTable) Select(npcType string) npcStatBlock {
	for _, block := range n {
		if strings.ToLower(block.Role) == strings.ToLower(npcType) {
			return block
		}
	}

	return n[len(n)-1]
}

// NPCStats are default values for NPCs of various roles
var NPCStats = npcStatBlockTable{
	{"Enforcer", 5, 3, 2, 2, "Intimidate 3, Fight 2, Force 1"},
	{"Gearhead", 2, 2, 5, 3, "Jury-Rig 3, Comprehend 2, Scout 1"},
	{"Stalker", 2, 5, 3, 2, "Find the Path 3, Shoot 2, Sneak 1"},
	{"Fixer", 2, 2, 3, 5, "Make a Deal 3, Manipulate 2, Move 1"},
	{"Dog Handler", 3, 4, 3, 2, "Sic a Dog 3, Shoot 2, Sneak 1"},
	{"Chronicler", 2, 2, 4, 4, "Inspire 3, Comprehend 2, Heal 1"},
	{"Boss", 3, 3, 2, 4, "Command 3, Shoot 2, Fight 1"},
	{"Slave", 4, 4, 2, 2, "Shake it Off 3, Endure 2, Fight 1"},
	{"No Role", 3, 3, 3, 3, "Level 2 in one skill"},
}

var npcGenHeaders = []string{"Name", "Appearance", "Goal", "Weapon"}

type npcGen struct {
	Name       roll.List
	Appearance roll.List
	Goal       roll.List
	Weapon     roll.List
}

// Enforcer table data to roll
var Enforcer = npcGen{
	Name: roll.List{
		Name:  "Name",
		Items: []string{"Hugust", "Ingrit", "lenny", "Marl", "Nelma", "Rebeth"},
	},
	Appearance: roll.List{
		Name: "Appearance",
		Items: []string{
			"Hulking posture, ape-like arms,grunts",
			"Short and stocky, hoarse voice",
			"Wiry, skinny and muscular",
			"Fat and with a wheezing voice",
			"Grotesquely tall, abnormal face",
			"Weathered and scarred, hissing voice",
		},
	},
	Goal: roll.List{
		Name: "Goal",
		Items: []string{
			"Stop fighting and build something instead",
			"Beat up anyone in her way",
			"Live another day, there is nothing else",
			"Torment others and hear them scream",
			"Please the Boss",
			"Show who’s strongest",
		},
	},
	Weapon: roll.List{
		Name: "Weapon",
		Items: []string{
			"Brass knuckles",
			"Bicycle chain",
			"Spiked bat",
			"Scrap knife",
			"Scrap axe",
			"Scrap pistol",
		},
	},
}

// Gearhead table data to roll
var Gearhead = npcGen{
	Name: roll.List{
		Name: "Name",
		Items: []string{
			"Quark",
			"Lambda",
			"Loranga",
			"Naphta",
			"Tetris",
			"Zippo",
		},
	},
	Appearance: roll.List{
		Name: "Appearance",
		Items: []string{
			"Skinny, spiked hair, always smiling",
			"Old blue coveralls, tools in every pocket",
			"Bald, extremely dirty, kind",
			"Burn-scarred skin, googles, smells of smoke",
			"Wears a headpiece made of scrap, mumbles constantly",
			"Wears a coat adorned by scrap",
		},
	},
	Goal: roll.List{
		Name: "Goal",
		Items: []string{
			"Find Eden and the truth about the People",
			"Recreate the technical marvels of the Old Age",
			"Build things to help others",
			"Blow stuff up",
			"Create a device that controls the minds of others",
			"Build a huge scrap sculpture in the Ark",
		},
	},
	Weapon: roll.List{
		Name: "Weapon",
		Items: []string{
			"Scrap pistol, Gear Bonus +2",
			"Scarp pistol with three barrels",
			"None",
			"Flamethrower",
			"-",
			"Bicycle chain",
		},
	},
}

// Stalker table to roll
var Stalker = npcGen{
	Name: roll.List{
		Name: "Name",
		Items: []string{
			"Danko",
			"Franton",
			"Hammed",
			"Jena",
			"Krin",
			"Tula",
		},
	},
	Appearance: roll.List{
		Name: "Appearance",
		Items: []string{
			"Always wears a gasmask, never speaks",
			"Scarred face, brags a lot",
			"Dead eyes, coughs from Rot damage, smells of booze",
			"Camouflage gear, constantly fiddles with her gun",
			"Skinny and wiry, always on her guard",
			"Wears a torn raincoat, pulls an old shopping cart",
		},
	},
	Goal: roll.List{
		Name: "Goal",
		Items: []string{
			"Go farther into the Zone than anyone else",
			"Be revered as a hero by the People",
			"Walk into the Zone and let the Rot take him",
			"Stalk a victim for days and then kill it",
			"Keep out of trouble, avoid others",
			"Find something worth fighting for",
		},
	},
	Weapon: roll.List{
		Name: "Weapon",
		Items: []string{
			"Scrap rifle, Gear Bonus +2",
			"Scrap pistol with two barrels",
			"Scrap spear",
			"Scrap rifle, weapon damage 3",
			"Bow, Gear Bonus +2",
			"Slingshot",
		},
	},
}

// Fixer table data to roll on
var Fixer = npcGen{
	Name: roll.List{
		Name: "Name",
		Items: []string{
			"Abed",
			"Denrik",
			"Fillix",
			"Jolisa",
			"Lula",
			"Monja",
		},
	},
	Appearance: roll.List{
		Name: "Appearance",
		Items: []string{
			"Bald and pudgy, constantly smiling",
			"Unnaturally handsome, shiny hair",
			"Stares at others’ gear, kleptomaniac",
			"Skinny, large staring eyes",
			"Abnormally short, talks constantly",
			"Wheezing voice, always wants to trade something",
		},
	},
	Goal: roll.List{
		Name: "Goal",
		Items: []string{
			"Make the deal of a lifetime",
			"Make others feel good and take advantage of it",
			"Collect a large stockpile of grub and gear",
			"Abnormally interested in artifacts and scrap",
			"Swindle the last bullet from everyone else",
			"Become a Boss one day",
		},
	},
	Weapon: roll.List{
		Name: "Weapon",
		Items: []string{
			"Scrap knife",
			"-",
			"Scrap pistol, Gear Bonus +2",
			"Bicycle chain",
			"Scrap pistol",
			"Scrap knife",
		},
	},
}
