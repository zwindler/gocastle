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
	LootXP    int
	LootGold  int
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
		LootXP:    100,
		LootGold:  0,
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
		LootXP:    500,
		LootGold:  100,
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

func (npc *NPCStats) HandleNPCDamage(damageDealt int) string {
	newHP := npc.CurrentHP - damageDealt

	// Here there are levels of injury
	// I want to give player additional information, but not every time!
	// only when NPC are going from above 80% live to under 80%, for example
	var additionalInfo string
	if newHP <= 0 {
		additionalInfo = fmt.Sprintf("%s is dead.", npc.Name)
	} else if newHP > 0 && newHP <= int(0.2*float64(npc.MaxHP)) && npc.CurrentHP > int(0.2*float64(npc.MaxHP)) {
		additionalInfo = fmt.Sprintf("%s looks barely alive.", npc.Name)
	} else if newHP > int(0.2*float64(npc.MaxHP)) && newHP <= int(0.5*float64(npc.MaxHP)) && npc.CurrentHP > int(0.5*float64(npc.MaxHP)) {
		additionalInfo = fmt.Sprintf("%s looks seriously injured.", npc.Name)
	} else if newHP > int(0.5*float64(npc.MaxHP)) && newHP <= int(0.8*float64(npc.MaxHP)) && npc.CurrentHP > int(0.8*float64(npc.MaxHP)) {
		additionalInfo = fmt.Sprintf("%s looks injured.", npc.Name)
	} else if newHP > int(0.8*float64(npc.MaxHP)) && newHP < npc.MaxHP && npc.CurrentHP == npc.MaxHP {
		additionalInfo = fmt.Sprintf("%s looks barely injured.", npc.Name)
	}
	return fmt.Sprintf("you strike at the %s, %s's hit! %s", npc.Name, npc.Pronoun, additionalInfo)
}

func (npc *NPCStats) IsNPCDead() bool {
	return npc.CurrentHP <= 0
}

// For a given NPCsOnCurrentMap, check all NPCs if one is located on x,y
func (NPCList *NPCsOnCurrentMap) GetNPCAtPosition(x, y int) int {
	// find if a NPC matches our destination
	for index, npc := range NPCList.List {
		if npc.Avatar.PosX == x && npc.Avatar.PosY == y {
			return index
		}
	}
	return -1
}

// For a given NPCsOnCurrentMap, remove NPC by list id and hide CanvasImage
func (NPCList *NPCsOnCurrentMap) RemoveNPCByIndex(index int) {
	// Check if the index is within the valid range of the slice.
	if index >= 0 && index < len(NPCList.List) {
		// Remove NPC image from map
		NPCList.List[index].Avatar.CanvasImage.Hidden = true
		// Use slicing to remove the element at the specified index.
		NPCList.List = append(NPCList.List[:index], NPCList.List[index+1:]...)
	}

}
