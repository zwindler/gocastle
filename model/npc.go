package model

import (
	"github.com/zwindler/gocastle/pkg/avatar"
	"github.com/zwindler/gocastle/pkg/hp"
	"github.com/zwindler/gocastle/pkg/mp"
	"github.com/zwindler/gocastle/pkg/npc"
)

var (
	FemaleFarmerAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/female-farmer.png",
	})
	FemaleFarmer = npc.New(npc.Stats{
		Name:    "Farmer",
		Avatar:  FemaleFarmerAvatar,
		Dialog:  "Hello, my name is Mylène :-)",
		Pronoun: "she",
		Hostile: false,
		HP:      hp.New(10),
	})

	FemaleMageAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/woman-mage.png",
	})
	FemaleMage = npc.New(npc.Stats{
		Name:    "Mage",
		Avatar:  FemaleMageAvatar,
		Pronoun: "she",
		Hostile: false,
		HP:      hp.New(15),
		MP:      mp.New(20),
	})

	KoboldAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/kobold-short.png",
	})
	Kobold = npc.New(npc.Stats{
		Name:     "Kobold",
		Avatar:   KoboldAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(4),
		LootXP:   30,
		LootGold: 2,
	})

	GoblinAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/goblin-short.png",
	})
	Goblin = npc.New(npc.Stats{
		Name:     "Goblin",
		Avatar:   GoblinAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(6),
		LootXP:   50,
		LootGold: 4,
	})

	GiantAntAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/giant-ant.png",
	})
	GiantAnt = npc.New(npc.Stats{
		Name:     "Giant Ant",
		Avatar:   GiantAntAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(10),
		LootXP:   60,
		LootGold: 0,
	})

	OrkAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/ork-short.png",
	})
	Ork = npc.New(npc.Stats{
		Name:     "Ork",
		Avatar:   OrkAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(14),
		LootXP:   80,
		LootGold: 10,
	})

	WolfAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/wolf.png",
	})
	Wolf = npc.New(npc.Stats{
		Name:     "Wolf",
		Avatar:   WolfAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(10),
		LootXP:   100,
		LootGold: 0,
	})

	GiantRedAntAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/giant-red-ant.png",
	})
	GiantRedAnt = npc.New(npc.Stats{
		Name:     "Giant Red Ant",
		Avatar:   GiantRedAntAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(20),
		LootXP:   150,
		LootGold: 0,
	})

	MimicAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/mimic.png",
	})
	Mimic = npc.New(npc.Stats{
		Name:     "Mimic",
		Avatar:   MimicAvatar,
		Pronoun:  "it",
		Hostile:  true,
		HP:       hp.New(25),
		LootXP:   300,
		LootGold: 500,
	})

	OgreAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/ogre.png",
	})
	Ogre = npc.New(npc.Stats{
		Name:     "Ogre",
		Avatar:   OgreAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(35),
		LootXP:   500,
		LootGold: 100,
	})

	MinotaurAvatar = avatar.New(avatar.Avatar{
		CanvasPath: "static/minotaur-short.png",
	})
	Minotaur = npc.New(npc.Stats{
		Name:     "Minotaur",
		Avatar:   MinotaurAvatar,
		Pronoun:  "he",
		Hostile:  true,
		HP:       hp.New(50),
		LootXP:   1000,
		LootGold: 300,
	})
)
