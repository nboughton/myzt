package table

import (
	"bytes"
	"fmt"
	"math/rand"
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

// Format npc output
func (n NPC) Format(t format.OutputType) string {
	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, format.Table(t, []string{},
		[][]string{
			{n.Name, n.Stats.Role},
			{"Appearance", n.Appearance},
			{"Goal", n.Goal},
			{"Weapon", n.Weapon},
			{"Mutation", Mutation.Roll()},
			{"MP", strconv.Itoa(rand.Intn(3) + 1)},
		}))
	fmt.Fprint(buf, n.Stats.Format(t))

	return buf.String()
}

// NewNPC roll an npcGen table
func NewNPC(npcType string) NPC {
	var n NPC

	if strings.ToLower(npcType) == "random" {
		r := NPCStats.Roles()
		npcType = r[rand.Intn(len(r))]
	}

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
	case "dog handler":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       DogHandler.Name.Roll(),
			Appearance: DogHandler.Appearance.Roll(),
			Goal:       DogHandler.Goal.Roll(),
			Weapon:     DogHandler.Weapon.Roll(),
		}
	case "chronicler":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Chronicler.Name.Roll(),
			Appearance: Chronicler.Appearance.Roll(),
			Goal:       Chronicler.Goal.Roll(),
			Weapon:     Chronicler.Weapon.Roll(),
		}
	case "boss":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Boss.Name.Roll(),
			Appearance: Boss.Appearance.Roll(),
			Goal:       Boss.Goal.Roll(),
			Weapon:     Boss.Weapon.Roll(),
		}
	case "slave":
		n = NPC{
			Stats:      NPCStats.Select(npcType),
			Name:       Slave.Name.Roll(),
			Appearance: Slave.Appearance.Roll(),
			Goal:       Slave.Goal.Roll(),
			Weapon:     Slave.Weapon.Roll(),
		}
	default:
		n = NPC{
			Stats: NPCStats.Select("No Role"),
			Name:  Names.Roll(),
		}
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

func (n npcStatBlockTable) Roles() []string {
	var out []string

	for _, r := range n {
		out = append(out, r.Role)
	}

	return out
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

// DogHandler table to roll on
var DogHandler = npcGen{
	Name: roll.List{
		Items: []string{
			"Anny",
			"Brie",
			"Finn",
			"Jony",
			"Krinnel",
			"Montiac",
		},
	},
	Appearance: roll.List{
		Items: []string{
			"Chews on a root, spits constantly",
			"Black bangs cover a deformed eye",
			"Greenish, rough skin, mostly drunk",
			"Face shaded by hood, silent and aggressive",
			"Only talks to her dog, never to people",
			"Crew cut, irritable and violent",
		},
	},
	Goal: roll.List{
		Items: []string{
			"Find someone other than her dog to care for",
			"Kill anyone who comments on her appearance",
			"Forget his sorrows",
			"Leave the Ark and find a better life",
			"Give her dog something good to eat",
			"Beat up anyone who speaks ill of his dog",
		},
	},
	Weapon: roll.List{
		Items: []string{
			"Slingshot",
			"Scrap knife",
			"Baseball bat",
			"Scrap axe",
			"Bow",
			"Brass knuckles",
		},
	},
}

// Chronicler table to roll on
var Chronicler = npcGen{
	Name: roll.List{
		Items: []string{
			"Astrina",
			"Danova",
			"Hanneth",
			"Maxim",
			"Silas",
			"Victon",
		},
	},
	Appearance: roll.List{
		Items: []string{
			"Red mohawk, sees good in everyone",
			"Wears a dress from the Old Age",
			"Black bangs, silent, takes notes constantly",
			"Skinny and weak, aged beyond his years",
			"Unnaturally pale, translucent skin, sad",
			"Boyish, reddish complexion, always happy",
		},
	},
	Goal: roll.List{
		Items: []string{
			"Find a hero who can save the People",
			"Understand what the Elder needs, and make it happen",
			"Rise from the misery and dedicate her life to study",
			"Write the story of how the People finds Eden",
			"To make the Ark evolve into an advanced society",
			"Become the hero in the story about the People",
		},
	},
	Weapon: roll.List{
		Items: []string{
			"Scrap knife",
			"-",
			"-",
			"-",
			"-",
			"Scrap knife",
		},
	},
}

// Boss table data to roll on
var Boss = npcGen{
	Name: roll.List{
		Items: []string{
			"Brictoria",
			"Gunitt",
			"Cristor",
			"Marlotte",
			"Mohammin",
			"Oscartion",
		},
	},
	Appearance: roll.List{
		Items: []string{
			"Cold eyes, scarred face, crew cut",
			"Short, skinny, wears a top hat",
			"Tall and skinny, fiery eyes",
			"Cool and detached",
			"Legless and is carried around by Slaves",
			"Extremely fat, can hardly move",
		},
	},
	Goal: roll.List{
		Items: []string{
			"Turn the People into an army and conquer the Zone",
			"Collect large amounts of artifacts",
			"Topple the Elder and establish equality in the Ark",
			"Develop the Ark and lead the People to a better life",
			"Rally more followers and leave the Ark with his cult",
			"Get rid of all rivals and take sole power in the Ark",
		},
	},
	Weapon: roll.List{
		Items: []string{
			"Scrap rifle, Gear Bonus +2",
			"Scrap pistol with three barrels",
			"Spiked bat",
			"-",
			"Scrap knife",
			"Scrap pistol, weapon damage 3",
		},
	},
}

// Slave table data to roll on
var Slave = npcGen{
	Name: roll.List{
		Items: []string{
			"Dink",
			"Eria",
			"Henny",
			"Hent",
			"Lin",
			"Wilo",
		},
	},
	Appearance: roll.List{
		Items: []string{
			"Silent and observant, bitter and treacherous",
			"Arms and legs like logs, grim",
			"Dirty, always hums on some melody",
			"Massive body with knobby skin, hunchback",
			"Wiry and grumpy, fights a lot",
			"Skinny and weak, a dreamer",
		},
	},
	Goal: roll.List{
		Items: []string{
			"Find someone to use to get free",
			"Toil away without complaining",
			"See the bright side of things",
			"Work and suffer until death sets him free",
			"To take her aggressions out on anyone",
			"To escape from the Ark and never come back",
		},
	},
	Weapon: roll.List{
		Items: []string{"-"},
	},
}

// Names to roll up
var Names = roll.List{
	Items: []string{
		"Abar", "Alia", "Alix", "Almenkir", "Amala",
		"Augustian", "Bark", "Benso", "Delta", "Doda", "Dunk",
		"Edin", "Elona", "Emalina", "Emdor", "Endel", "Ergom",
		"Erian", "Erister", "Felin", "Fils", "Gerber", "Grits",
		"Gros", "Hild", "Ibit", "Iridia", "Johalin", "Jol", "Jonar",
		"Juna", "Juperia", "Kalgu", "Kaska", "Katin", "Kim",
		"Koli", "Leodor", "Landon", "Lard", "Linna", "Margot",
		"Marlian", "Makron", "Markot", "Mart", "Max",
		"Mintel", "Miri", "Mohan", "Molger", "Morris",
		"Mubba", "Natara", "Nelma", "Nely", "Novia", "Octane",
		"Olias", "Olli", "Oslo", "Otiak", "Plonk", "Piro", "Pontis",
		"Rasper", "Regis", "Rog", "Rutger", "Sagabet", "Sari",
		"Savik", "Silia", "Siniss", "Slabados", "Sofin", "Styx",
		"Tegra", "Torry", "Vabb", "Vang", "Vilip", "Visek",
		"Vorhan", "Zanova", "Zeb", "Zeke", "Zingo",
	},
}
