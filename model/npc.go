package model

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zwindler/gocastle/pkg/hp"
)

type NPCStats struct {
	Name      string
	Pronoun   string
	Dialog    string
	Hostile   bool
	Avatar    Avatar
	HP        *hp.HP
	MaxMP     int
	CurrentMP int
	LootXP    int
	LootGold  int
}

var (
	FemaleFarmerAvatar = Avatar{
		CanvasPath: "static/female-farmer.png",
	}
	FemaleFarmer = NPCStats{
		Name:    "Farmer",
		Avatar:  FemaleFarmerAvatar,
		Dialog:  "Hello, my name is Mylène :-)",
		Pronoun: "she",
		Hostile: false,
		HP:      hp.New(10),
	}

	FemaleMageAvatar = Avatar{
		CanvasPath: "static/woman-mage.png",
	}
	FemaleMage = NPCStats{
		Name:    "Mage",
		Avatar:  FemaleMageAvatar,
		Pronoun: "she",
		Hostile: false,
		HP:      hp.New(15),
		MaxMP:   20,
	}

	KoboldAvatar = Avatar{
		CanvasPath: "static/kobold-short.png",
	}
	Kobold = NPCStats{
		Name:     "Kobold",
		Avatar:   KoboldAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(4),
		LootXP:   30,
		LootGold: 2,
	}

	GoblinAvatar = Avatar{
		CanvasPath: "static/goblin-short.png",
	}
	Goblin = NPCStats{
		Name:     "Goblin",
		Avatar:   GoblinAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(6),
		LootXP:   50,
		LootGold: 4,
	}

	GiantAntAvatar = Avatar{
		CanvasPath: "static/giant-ant.png",
	}
	GiantAnt = NPCStats{
		Name:     "Giant Ant",
		Avatar:   GiantAntAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(10),
		LootXP:   60,
		LootGold: 0,
	}

	OrkAvatar = Avatar{
		CanvasPath: "static/ork-short.png",
	}
	Ork = NPCStats{
		Name:     "Ork",
		Avatar:   OrkAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(14),
		LootXP:   80,
		LootGold: 10,
	}

	WolfAvatar = Avatar{
		CanvasPath: "static/wolf.png",
	}
	Wolf = NPCStats{
		Name:     "Wolf",
		Avatar:   WolfAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(10),
		LootXP:   100,
		LootGold: 0,
	}

	GiantRedAntAvatar = Avatar{
		CanvasPath: "static/giant-red-ant.png",
	}
	GiantRedAnt = NPCStats{
		Name:     "Giant Red Ant",
		Avatar:   GiantRedAntAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(20),
		LootXP:   150,
		LootGold: 0,
	}

	MimicAvatar = Avatar{
		CanvasPath: "static/mimic.png",
	}
	Mimic = NPCStats{
		Name:     "Mimic",
		Avatar:   MimicAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(25),
		LootXP:   300,
		LootGold: 500,
	}

	OgreAvatar = Avatar{
		CanvasPath: "static/ogre.png",
	}
	Ogre = NPCStats{
		Name:     "Ogre",
		Avatar:   OgreAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(35),
		LootXP:   500,
		LootGold: 100,
	}

	MinotaurAvatar = Avatar{
		CanvasPath: "static/minotaur-short.png",
	}
	Minotaur = NPCStats{
		Name:     "Minotaur",
		Avatar:   MinotaurAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(50),
		LootXP:   1000,
		LootGold: 300,
	}
)

// CreateNPC creates a copy of a given NPC at given coordinates.
func CreateNPC(npc NPCStats, coord Coord) *NPCStats {
	avatar := CreateAvatar(npc.Avatar, coord)
	return &NPCStats{
		Name:    npc.Name,
		Pronoun: npc.Pronoun,
		Avatar:  avatar,
		Dialog:  npc.Dialog,
		Hostile: npc.Hostile,
		HP:      npc.HP,
		MaxMP:   npc.MaxMP,
		CurrentMP: func() int {
			if npc.CurrentMP == 0 {
				return npc.MaxMP
			}
			return npc.CurrentMP
		}(),
		LootXP:   npc.LootXP,
		LootGold: randomizeGoldLoot(npc.LootGold),
	}
}

// HandleNPCDamage returns strings for having nice logs during combat with NPCs.
func (npc *NPCStats) HandleNPCDamage() string {
	var additionalInfo string

	// Here there are levels of injury
	// I want to give player additional information, but not every time!
	// only when NPC are going from above 80% live to under 80%, for example
	switch {
	case npc.HP.IsDead():
		additionalInfo += fmt.Sprintf("%s is dead.", npc.Name)
	// remaininghp between 80% and 100%
	case npc.HP.IsAlive() && npc.HP.Percent() > 0.8:
		additionalInfo += fmt.Sprintf("%s looks barely injured.", npc.Name)
	// remaininghp between 50% and 80%
	case npc.HP.IsAlive() && npc.HP.Percent() > 0.5:
		additionalInfo += fmt.Sprintf("%s looks injured.", npc.Name)
	// remaininghp between 20% and 50%
	case npc.HP.IsAlive() && npc.HP.Percent() > 0.2:
		additionalInfo += fmt.Sprintf("%s looks seriously injured.", npc.Name)
	// remaininghp between 0% and 20%
	case npc.HP.IsAlive() && npc.HP.Percent() > 0:
		additionalInfo += fmt.Sprintf("%s looks barely alive.", npc.Name)
	}
	return fmt.Sprintf("you strike at the %s, %s's hit! %s", npc.Name, npc.Pronoun, additionalInfo)
}

// IsNPCDead checks if NPC's HP <= 0.
func (npc *NPCStats) IsNPCDead() bool {
	return npc.HP.IsDead()
}

// randomizeGoldLoot generates a random amount of gold within a specified range.
func randomizeGoldLoot(goldAmount int) int {
	if goldAmount <= 0 {
		return 0
	}

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random multiplier between 0.5 and 1.5 (inclusive)
	multiplier := rand.Float64() + 0.5 //nolint:gosec

	// Calculate the randomized gold amount
	randomizedGold := int(float64(goldAmount) * multiplier)

	return randomizedGold
}
