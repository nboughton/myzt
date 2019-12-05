package table

import "github.com/nboughton/go-roll"

// SectorEnv ironment table
var SectorEnv = roll.Table{
	Name: "Sector Environment",
	ID:   "sector.environment",
	Dice: roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12}, Text: "Thick Woods", Action: func() string {
			return ""
		}},
	},
}

// RuinsNormal table
var RuinsNormal = roll.List{
	Items: []string{
		"Airplane Wreck",
		"Amusement Park",
		"Battlefield",
		"Bus Station",
		"Car Park ",
		"Church ",
		"Cinema ",
		"Crater ",
		"Dilapidated Mansion ",
		"Fast Food Joint ",
		"Gas Station ",
		"Highway ",
		"Hospital",
		"Hunting Store",
		"Mall",
		"Marina",
		"Museum",
		"Office Building",
		"Overgrown Park",
		"Playground",
		"Police Station",
		"Radio Station",
		"Residential Blocks",
		"Road Tunnel",
		"Ruined Bridge",
		"School",
		"Shelter",
		"Skyscraper",
		"Sports Center",
		"Suburbia",
		"Subway Station",
		"Supermarket",
		"Swimming Hall",
		"Tank",
		"Theater",
		"Train Station",
	},
}

// RuinsIndustrial table
var RuinsIndustrial = roll.List{
	Items: []string{
		"Factory",
		"Military Base",
		"Oil Cistern",
		"Pipeline",
		"Purification Plant",
		"Power Line",
		"Radio Mast",
		"Refinery",
		"Rubbish Dump",
		"Shipwreck",
		"Shooting Range",
		"Windmill",
	},
}

// RotLevel table
var RotLevel = roll.Table{
	Name: "Rot Level",
	ID:   "sector.rotLevel",
	Dice: roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12}, Text: "0 - Rot Oasis. The PCs are safe from the Rot here."},
		{Match: []int{13, 14, 15, 16, 21, 22, 23, 24, 25, 26, 31, 32, 33, 34, 35, 36, 41, 42, 43, 44, 45, 46, 51, 52, 53, 54, 55}, Text: "1 - Weak Rot. The PCs suffer one Rot Point per day spent in such sectors."},
		{Match: []int{56, 61, 62, 63, 64, 65, 66}, Text: "2 - Rot-Heavy Area. The PCs suffer one Rot Point per hour."},
	},
}

// ThreatType table
var ThreatType = roll.Table{
	Name: "Threat Type",
	ID:   "sector.threat.type",
	Dice: roll.Dice{N: 1, Die: roll.D6},
	Items: []roll.TableItem{
		{Match: []int{1, 2}, Text: "Humanoid"},
		{Match: []int{3, 4, 5}, Text: "Monster"},
		{Match: []int{6}, Text: "Phenomenon"},
	},
}

// ThreatHumanoid table
var ThreatHumanoid = roll.Table{
	Name: "Humanoid Thread",
	ID:   "sector.threat.humanoid",
	Dice: roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11}, Text: "Amnesiac"},
		{Match: []int{12, 13}, Text: "Beast Mutants"},
		{Match: []int{14, 15, 16}, Text: "Cannibals"},
		{Match: []int{21, 22}, Text: "Doom Cultists"},
		{Match: []int{23, 24, 25}, Text: "Exiled Mutants"},
		{Match: []int{26, 31, 32}, Text: "Expedition from another Ark"},
		{Match: []int{33, 34, 35}, Text: "Helldrivers"},
		{Match: []int{36, 41, 42}, Text: "Morlocks"},
		{Match: []int{43, 44}, Text: "Nova Cultists"},
		{Match: []int{45, 46}, Text: "Patrol from the Ark"},
		{Match: []int{51, 52}, Text: "Scrap Oracle"},
		{Match: []int{53, 54}, Text: "Wanderers"},
		{Match: []int{55, 56}, Text: "Water Trader"},
		{Match: []int{61, 62}, Text: "Wreckers"},
		{Match: []int{63, 64, 65, 66}, Text: "Zone-Ghouls"},
	},
}

// ThreatMonster table
var ThreatMonster = roll.Table{
	Name: "Monster Threat",
	ID:   "sector.threat.monster",
	Dice: roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12}, Text: "Acid Grass"},
		{Match: []int{13}, Text: "Air Jellies"},
		{Match: []int{14}, Text: "Automaton"},
		{Match: []int{15, 16}, Text: "Bitterbeasts"},
		{Match: []int{21, 22}, Text: "Deathswarm"},
		{Match: []int{23, 24}, Text: "Devourer"},
		{Match: []int{25, 26}, Text: "Grazers"},
		{Match: []int{31}, Text: "Gutfish (infected water)"},
		{Match: []int{32}, Text: "Killer Tree"},
		{Match: []int{33}, Text: "Mind Mosquitos"},
		{Match: []int{34}, Text: "Nightmare Flowers"},
		{Match: []int{35}, Text: "Parasite Fungus"},
		{Match: []int{36}, Text: "Razorback"},
		{Match: []int{41, 42}, Text: "Rot Ants"},
		{Match: []int{43}, Text: "Rotfish"},
		{Match: []int{44, 45}, Text: "Scrap Crows"},
		{Match: []int{46}, Text: "Trash Hawk"},
		{Match: []int{51}, Text: "Worm Swarm"},
		{Match: []int{52, 53, 54}, Text: "Zone Dogs"},
		{Match: []int{55, 56}, Text: "Zone Leeches"},
		{Match: []int{61, 62, 63}, Text: "Zone Rats"},
		{Match: []int{64, 65}, Text: "Zone Spider"},
		{Match: []int{66}, Text: "Zone Wasp"},
	},
}

// ThreatPhenomena table
var ThreatPhenomena = roll.Table{
	Name:  "Phenomena Threat",
	ID:    "sector.threat.phenom",
	Dice:  roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{},
}
