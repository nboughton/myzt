package table

import (
	"fmt"

	"github.com/nboughton/go-roll"
)

const (
	nrna = "%s\nNo Ruins\n%s\nNo Artifact\n%s"
	ra   = "%s\n%s\n%s\n%s\n%s"
)

// SectorEnv ironment table
var SectorEnv = roll.Table{
	Name: "Sector Environment",
	ID:   "sector.environment",
	Dice: roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12}, Text: fmt.Sprintf(nrna, "Thick Woods", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{13, 14, 15}, Text: fmt.Sprintf(nrna, "Scrublands", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{16, 21}, Text: fmt.Sprintf(nrna, "Marshlands", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{22, 23, 24}, Text: fmt.Sprintf(nrna, "Dead Woods", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{25, 26}, Text: fmt.Sprintf(nrna, "Ash Desert", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{31}, Text: fmt.Sprintf(nrna, "Huge Crater", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{32}, Text: fmt.Sprintf(nrna, "Glassified Field", ThreatType.Roll(), RotLevel.Roll())},
		{Match: []int{33, 34, 35}, Text: fmt.Sprintf(ra, "Overgrown Ruins", RuinsNormal.Roll(), ThreatType.Roll(), Artifact.Roll(), RotLevel.Roll())},
		{Match: []int{36, 41, 42}, Text: fmt.Sprintf(ra, "Crumbling Ruins", RuinsNormal.Roll(), ThreatType.Roll(), Artifact.Roll(), RotLevel.Roll())},
		{Match: []int{43, 44, 45, 46, 51}, Text: fmt.Sprintf(ra, "Decayed Ruins", RuinsNormal.Roll(), ThreatType.Roll(), Artifact.Roll(), RotLevel.Roll())},
		{Match: []int{52, 53, 54, 55, 56}, Text: fmt.Sprintf(ra, "Unscathed Ruins", RuinsNormal.Roll(), ThreatType.Roll(), Artifact.Roll(), RotLevel.Roll())},
		{Match: []int{61, 62, 63, 64}, Text: fmt.Sprintf(ra, "Derelict Industries", RuinsIndustrial.Roll(), ThreatType.Roll(), Artifact.Roll(), RotLevel.Roll())},
		{Match: []int{65, 66}, Text: "Settlement"},
	},
}

// RuinsNormal table
var RuinsNormal = roll.List{
	Items: []string{
		"Airplane Wreck",
		"Amusement Park",
		"Battlefield",
		"Bus Station",
		"Car Park",
		"Church",
		"Cinema",
		"Crater",
		"Dilapidated Mansion",
		"Fast Food Joint",
		"Gas Station",
		"Highway",
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
		{Match: []int{1, 2}, Text: ThreatHumanoid.Roll()},
		{Match: []int{3, 4, 5}, Text: ThreatMonster.Roll()},
		{Match: []int{6}, Text: ThreatPhenomena.Roll()},
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
	Name: "Phenomena Threat",
	ID:   "sector.threat.phenom",
	Dice: roll.Dice{N: 1, Die: roll.D66},
	Items: []roll.TableItem{
		{Match: []int{11, 12, 13}, Text: "Acid Rain"},
		{Match: []int{14, 15}, Text: "Ash Storm"},
		{Match: []int{16, 21}, Text: "Dust Tornado"},
		{Match: []int{22, 23}, Text: "Electric Storm"},
		{Match: []int{24}, Text: "Ghost Lights"},
		{Match: []int{25}, Text: "Inertia Field"},
		{Match: []int{26, 31}, Text: "Magnetic Field"},
		{Match: []int{31}, Text: "Mirage"},
		{Match: []int{33, 34}, Text: "Mud Puddles"},
		{Match: []int{35, 36}, Text: "Night Lights"},
		{Match: []int{41}, Text: "Obelisk"},
		{Match: []int{42, 43}, Text: "Pillars of Light"},
		{Match: []int{44, 45, 46}, Text: "Rot Hotspot"},
		{Match: []int{51, 52}, Text: "Rot Wind"},
		{Match: []int{53, 54}, Text: "Sinkhole"},
		{Match: []int{55, 56}, Text: "Temperature Drop / Heat Wave"},
		{Match: []int{61, 62}, Text: "Unexploded Ordnance"},
		{Match: []int{63}, Text: "Vacuum"},
		{Match: []int{64, 65, 66}, Text: "Zone Smog"},
	},
}

// Artifact table
var Artifact = roll.List{
	Items: []string{
		"Air Mattress",
		"Antidepressants",
		"Assault Rifle",
		"Automobile",
		"Battery",
		"Bicycle",
		"Binoculars",
		"Bow & Arrow",
		"Canoe",
		"Cassette Player (Metaplot)",
		"Chainsaw",
		"Comic Book",
		"Compass",
		"Crossbow",
		"Diary",
		"Diving Gear",
		"Dress",
		"Energy Pills",
		"Flashlight",
		"Flare Gun",
		"Gas Mask",
		"Generator",
		"Guitar",
		"Hand Grenade",
		"Hockey Helmet",
		"Hunting Rifle",
		"ID Card (Metaplot)",
		"Jerrycan",
		"Katana",
		"Kevlar Vest",
		"Lifestyle Magazine",
		"Map of the Zone",
		"Motor Boat",
		"Painkillers",
		"Painting",
		"Perfume Bottle",
		"Protective Suit",
		"REGEN",
		"Revolver",
		"Scooter",
		"Semi-Automatic Pistol",
		"Shotgun",
		"Smoke Grenade",
		"Soda Can",
		"Stimulants",
		"Sunglasses",
		"Tuxedo",
		"Umbrella",
		"Video Camera (Metaplot)",
		"Wrench",
	},
}
