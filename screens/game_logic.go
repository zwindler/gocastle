package screens

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2/dialog"

	"github.com/zwindler/gocastle/maps"
	"github.com/zwindler/gocastle/pkg/timespent"
)

// actOnDirectionKey take player's new coordinates and act on it.
func actOnDirectionKey(newX, newY int) {
	// before doing anything, check if we aren't out of bounds
	if currentMap.CheckOutOfBounds(newX, newY) {
		// Player tries to escape map, prevent this, lose 2 seconds
		addLogEntry("you are blocked!")
		timespent.Increment(2)
	} else {
		// let's check if we find a NPC on our path
		if npc := currentMap.GetNPCAtPosition(newX, newY); npc != nil {
			// yes, but is the NPC hostile?
			if npc.Hostile {
				// let's attack!
				// TODO add some randomization
				npc.HP.Damage(player.PhysicalDamage)
				addLogEntry(npc.HandleNPCDamage())
				if npc.IsNPCDead() {
					if player.ChangeXP(npc.LootXP) {
						levelUpEntry := fmt.Sprintf("Level up! You are now level %d", player.Level)
						addLogEntry(levelUpEntry)
						levelUpPopup := showLevelUpScreen()
						dialog.ShowCustomConfirm("Level up!", "Validate", "Close", levelUpPopup, func(validate bool) {
							player.RefreshStats(true)
							updateStatsArea()
						}, currentWindow)
					}
					player.ChangeGold(npc.LootGold)
					currentMap.RemoveNPC(npc)
				}
				// attacking costs 5 seconds
				timespent.Increment(5)
			} else {
				// NPC is not hostile, we don't want to hurt them, but lost 2s
				if npc.Dialog != "" {
					dialogEntry := fmt.Sprintf("%s says: %s", npc.Name, npc.Dialog)
					addLogEntry(dialogEntry)
				} else {
					blockEntry := fmt.Sprintf("%s is blocking you", npc.Name)
					addLogEntry(blockEntry)
				}

				timespent.Increment(2)
			}
		} else {
			// no NPC found on our path, let's check if we can move
			if currentMap.CheckTileIsWalkable(newX, newY) {
				// path is free, let's move (3sec cost)
				player.Avatar.MoveAvatar(mapContainer, newX, newY)
				timespent.Increment(3)

				// this tile could be special, check if it is
				tile := currentMap.CheckTileIsSpecial(newX, newY)
				if tile != maps.NotSpecialTile {
					if tile.Type == "MapTransition" {
						currentMap = &maps.AllTheMaps[tile.Destination.Map]
						player.Avatar.Coord = tile.Destination
						ShowGameScreen(currentWindow)
					}
					// TODO handle error
				}
			} else {
				// you "hit" a wall, but lost 2s
				addLogEntry("you are blocked!")
				timespent.Increment(2)
			}
		}
	}
}

// newTurnForNPCs manages all the map's NPCs actions.
func newTurnForNPCs() {
	// for all NPCs, move
	for _, npc := range currentMap.NPCList {
		var newX, newY int
		if npc.Hostile && npc.Avatar.DistanceFromAvatar(&player.Avatar) <= 10 {
			// player is near, move toward him/her
			newX, newY = npc.Avatar.MoveAvatarTowardsAvatar(&player.Avatar)
		} else {
			// move randomly
			newX = npc.Avatar.Coord.X + rand.Intn(3) - 1 //nolint:gosec
			newY = npc.Avatar.Coord.Y + rand.Intn(3) - 1 //nolint:gosec
		}

		// don't check / try to move if coordinates stay the same
		if newX != npc.Avatar.Coord.X || newY != npc.Avatar.Coord.Y {
			// before doing anything, check if we aren't out of bounds
			if !currentMap.CheckOutOfBounds(newX, newY) {
				// let's check if we find another NPC on our NPC's path
				if otherNPC := currentMap.GetNPCAtPosition(newX, newY); otherNPC != nil {
					if (npc.Hostile && !otherNPC.Hostile) ||
						(!npc.Hostile && otherNPC.Hostile) {
						// TODO hostile NPC should attack friendly NPC
						// and vice versa
						addLogEntry(fmt.Sprintf("%s tries to attack %s", npc.Name, otherNPC.Name))
					}
					// let's then check we don't collide with player
				} else if player.Avatar.CollideWithPlayer(newX, newY) {
					if npc.Hostile {
						// TODO hostile NPC should attack player
						addLogEntry(fmt.Sprintf("%s tries to attack you", npc.Name))
					}
					// no ones in our NPC's way
				} else if currentMap.CheckTileIsWalkable(newX, newY) {
					npc.Avatar.MoveAvatar(mapContainer, newX, newY)
				}
			}
		}
	}
}
