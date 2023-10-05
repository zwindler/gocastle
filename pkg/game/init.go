package game

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"

	"github.com/zwindler/gocastle/pkg/avatar"
	"github.com/zwindler/gocastle/pkg/character"
	"github.com/zwindler/gocastle/pkg/coord"
	"github.com/zwindler/gocastle/pkg/embedmaps"
	"github.com/zwindler/gocastle/pkg/hp"
	"github.com/zwindler/gocastle/pkg/maps"
	"github.com/zwindler/gocastle/pkg/mp"
	"github.com/zwindler/gocastle/pkg/npc"
	"github.com/zwindler/gocastle/pkg/object"
)

var (
	PlayerAvatar = avatar.Avatar{}
	Player       = character.Stats{
		// temporary, for dev
		CharacterName: "zwindler",
		PointsToSpend: 0,
		// end temporary for dev
		Avatar: PlayerAvatar,
		// PointsToSpend:     10,
		StrengthValue:     10,
		ConstitutionValue: 10,
		IntelligenceValue: 10,
		DexterityValue:    10,
		Level:             1,
		HP:                hp.New(character.BaseHP),
		MP:                mp.New(character.BaseMP),
	}
	CurrentMap = &maps.AllTheMaps[0]
)

// InitGame will initialise all needed variables before start game (start=true) or load game (start=false).
func InitGame(window fyne.Window, start bool) {
	// refresh player stats (heal or not depending on "start")
	Player.RefreshStats(start)

	// init categories
	object.InitializeCategories()

	// create player Avatar
	if start {
		// load all pregenerated maps from json
		for i := 0; i < len(maps.AllTheMaps); i++ {
			thisMapMatrix, err := embedmaps.GetMapMatrixFromEmbed(fmt.Sprintf("maps/%d.json", i))
			if err != nil {
				dialog.ShowError(err, window)
			}
			maps.AllTheMaps[i].MapMatrix = thisMapMatrix
		}

		Player.ChangeGold(10)

		// TODO rework this
		// Map0 Village
		knife, _ := object.CreateObject(object.HuntingKnife, coord.Coord{X: 10, Y: 10, Map: 0})
		maps.AllTheMaps[0].ObjectList = append(maps.AllTheMaps[0].ObjectList, &knife)
		farmer := npc.Spawn(npc.FemaleFarmer, coord.Coord{X: 10, Y: 15, Map: 0})
		ant1 := npc.Spawn(npc.GiantAnt, coord.Coord{X: 5, Y: 34, Map: 0})
		ant2 := npc.Spawn(npc.GiantAnt, coord.Coord{X: 6, Y: 32, Map: 0})
		ant3 := npc.Spawn(npc.GiantAnt, coord.Coord{X: 7, Y: 33, Map: 0})
		maps.AllTheMaps[0].NPCList = append(maps.AllTheMaps[0].NPCList, farmer, ant1, ant2, ant3)

		// Map1 To The Old Mine
		sword, _ := object.CreateObject(object.BluntSword, coord.Coord{X: 9, Y: 4, Map: 1})
		maps.AllTheMaps[1].ObjectList = append(maps.AllTheMaps[1].ObjectList, &sword)
		wolf1 := npc.Spawn(npc.Wolf, coord.Coord{X: 70, Y: 24, Map: 1})
		wolf2 := npc.Spawn(npc.Wolf, coord.Coord{X: 69, Y: 23, Map: 1})
		maps.AllTheMaps[1].NPCList = append(maps.AllTheMaps[1].NPCList, wolf1, wolf2)

		Player.Avatar.Coord = coord.Coord{X: 15, Y: 15, Map: 0}
	}

	// pregenerate the map image to save time in game screen
	CurrentMap.GenerateMapImage()

	CurrentMap = &maps.AllTheMaps[Player.Avatar.Coord.Map]
	Player.Avatar = avatar.Spawn(Player.Avatar, Player.Avatar.Coord)
}
