package model

import (
	"fmt"

	"fyne.io/fyne/v2/canvas"
)

type NPCStats struct {
	Name      string
	Pronoun   string
	Hostile   bool
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
		Pronoun:   "he",
		Hostile:   false,
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
		Hostile:   true,
		MaxHP:     10,
		CurrentHP: 10,
	}

	MageAvatar = Avatar{
		CanvasPath: "./static/woman-mage.png",
	}
	Mage = NPCStats{
		Name:      "mage",
		Avatar:    MageAvatar,
		Pronoun:   "she",
		Hostile:   false,
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
		Pronoun:   "he",
		Hostile:   true,
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
		Hostile:   npc.Hostile,
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

func handleNPCDamage(npc *NPCStats, damageDealt int) string {
	newHP := npc.CurrentHP - damageDealt

	// Here there are levels of injury
	// I want to give player additionnal information, but not everytime!
	// only when NPC are going from above 80% live to under 80%, for example
	var additionnalInfo string
	if newHP <= 0 {
		additionnalInfo = fmt.Sprintf("%s is dead.", npc.Name)
	} else if newHP > 0 && newHP <= int(0.2*float64(npc.MaxHP)) && npc.CurrentHP > int(0.2*float64(npc.MaxHP)) {
		additionnalInfo = fmt.Sprintf("%s looks barely alive.", npc.Name)
	} else if newHP > int(0.2*float64(npc.MaxHP)) && newHP <= int(0.5*float64(npc.MaxHP)) && npc.CurrentHP > int(0.5*float64(npc.MaxHP)) {
		additionnalInfo = fmt.Sprintf("%s looks seriously injured.", npc.Name)
	} else if newHP > int(0.5*float64(npc.MaxHP)) && newHP <= int(0.8*float64(npc.MaxHP)) && npc.CurrentHP > int(0.8*float64(npc.MaxHP)) {
		additionnalInfo = fmt.Sprintf("%s looks injured.", npc.Name)
	} else if newHP > int(0.8*float64(npc.MaxHP)) && newHP < npc.MaxHP && npc.CurrentHP == npc.MaxHP {
		additionnalInfo = fmt.Sprintf("%s looks barely injured.", npc.Name)
	}
	return fmt.Sprintf("you strike at the %s, %s is hit! %s", npc.Name, npc.Pronoun, additionnalInfo)
}
