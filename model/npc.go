package model

import "fyne.io/fyne/v2/canvas"

type NPCStats struct {
	Name      string
	Pronoun   string
	Avatar    Avatar
	MaxHP     int
	CurrentHP int
	MaxMP     int
	CurrentMP int
}
type NPCsOnCurrentMap struct {
	List []NPCStats
}

var (
	FarmerAvatar = Avatar{
		CanvasPath: "./static/male-farmer.png",
	}
	Farmer = NPCStats{
		Name:      "farmer",
		Avatar:    FarmerAvatar,
		Pronoun:   "him",
		MaxHP:     10,
		CurrentHP: 10,
	}

	WolfAvatar = Avatar{
		CanvasPath: "./static/wolf.png",
	}
	Wolf = NPCStats{
		Name:      "wolf",
		Avatar:    WolfAvatar,
		Pronoun:   "it",
		MaxHP:     10,
		CurrentHP: 10,
	}

	MageAvatar = Avatar{
		CanvasPath: "./static/woman-mage.png",
	}
	
	Mage = NPCStats{
		Name:      "mage",
		Avatar:    MageAvatar,
		Pronoun:   "her",
		MaxHP:     15,
		CurrentHP: 15,
		MaxMP:     20,
		CurrentMP: 20,
	}

	OgreAvatar = Avatar{
		CanvasPath: "./static/ogre.png",
	}
	
	Ogre = NPCStats{
		Name:      "ogre",
		Avatar:    OgreAvatar,
		Pronoun:   "him",
		MaxHP:     25,
		CurrentHP: 25,
		MaxMP:     0,
		CurrentMP: 0,
	}
	
)

func CreateNPC(npc NPCStats, x, y int) NPCStats {
	avatar := createAvatar(npc.Avatar, x, y)
	return NPCStats{
		Name:      npc.Name,
		Pronoun:   npc.Pronoun,
		Avatar:    avatar,
		MaxHP:     npc.MaxHP,
		CurrentHP: npc.CurrentHP,
		MaxMP:     npc.MaxMP,
		CurrentMP: npc.CurrentMP,
	}
}

func createAvatar(avatar Avatar, x, y int) Avatar {
	return Avatar{
		CanvasImage: canvas.NewImageFromFile(avatar.CanvasPath),
		PosX:        x,
		PosY:        y,
	}
}
