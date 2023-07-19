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
		CanvasPath: "./static/farmer.png",
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
