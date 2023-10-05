package character_test

import (
	"testing"

	"github.com/zwindler/gocastle/pkg/character"
	"github.com/zwindler/gocastle/pkg/hp"
	"github.com/zwindler/gocastle/pkg/mp"
	"github.com/zwindler/gocastle/pkg/object"
	"github.com/zwindler/gocastle/pkg/pts"
)

func TestDeterminePhysicalDamage(t *testing.T) {
	player := &character.Stats{
		StrengthValue:  10,
		DexterityValue: 10,
		Inventory: []*object.Object{
			{
				Stats: []object.ObjectStat{
					{Name: "physicalDamage", Modifier: 5},
				},
				Equipped: true,
			},
		},
	}

	player.DeterminePhysicalDamage()

	if player.PhysicalDamage != 7 { // BasePhysicalDamage + 5 from the equipped item
		t.Errorf("Expected PhysicalDamage to be 7, but got %d", player.PhysicalDamage)
	}
}

func TestChangeXP(t *testing.T) {
	player := &character.Stats{
		Level: 1,
		HP:    hp.New(character.BaseHP),
		MP:    mp.New(character.BaseMP),
	}

	player.ChangeXP(100)

	if player.CurrentXP != 100 {
		t.Errorf("Expected CurrentXP to be 100, got %d", player.CurrentXP)
	}
}

func TestChangeGold(t *testing.T) {
	player := &character.Stats{}
	player.ChangeGold(10)

	if player.CurrentGold != 10 {
		t.Errorf("Expected player's gold to be 10, but got %d", player.CurrentGold)
	}
}

func TestDetermineLevel(t *testing.T) {
	tcs := []struct {
		player          character.Stats
		expectedLevel   int
		expectedLevelUp bool
	}{
		{character.Stats{
			Level:     1,
			HP:        hp.New(character.BaseHP),
			MP:        mp.New(character.BaseMP),
			CurrentXP: 50,
		}, 1, false},
		{character.Stats{
			Level:     1,
			HP:        hp.New(character.BaseHP),
			MP:        mp.New(character.BaseMP),
			CurrentXP: 150,
		}, 2, true},
		{character.Stats{
			Level:     1,
			HP:        hp.New(character.BaseHP),
			MP:        mp.New(character.BaseMP),
			CurrentXP: 350,
		}, 3, true},
	}

	for _, tc := range tcs {
		leveledUp := tc.player.DetermineLevel()
		if leveledUp {
			if !tc.expectedLevelUp {
				t.Errorf("Didn't expect player to level up, but did")
			}
		} else {
			if tc.expectedLevelUp {
				t.Errorf("Expected player to level up, but didn't")
			}
		}

		if tc.player.Level != tc.expectedLevel {
			t.Errorf("Expected player level to be %d, got %d", tc.expectedLevel, tc.player.Level)
		}
	}
}

func TestRefreshStats(t *testing.T) {
	tcHP := hp.New(character.BaseHP)
	tcHP.Damage(5)

	tcs := []struct {
		player        character.Stats
		heal          bool
		expectedHP    *pts.Point
		expectedMaxHP *pts.Point
		expectedMP    *pts.Point
		expectedMaxMP *pts.Point
	}{
		{character.Stats{
			Level:             1,
			ConstitutionValue: 19, // 3 bonus points (19-10)/3*level
			IntelligenceValue: 10, // no bonus points
			HP:                hp.New(character.BaseHP),
			MP:                mp.New(character.BaseMP),
		}, false, pts.New(8), pts.New(11), pts.New(8), pts.New(8)},
		{character.Stats{
			Level:             1,
			ConstitutionValue: 10,   // no bonus points
			IntelligenceValue: 19,   // 3 bonus points (19-10)/3*level
			HP:                tcHP, // 8 HP - 5 damage
			MP:                mp.New(character.BaseMP),
		}, false, pts.New(3), pts.New(8), pts.New(8), pts.New(11)},
		{character.Stats{
			Level:             1,
			ConstitutionValue: 10,   // no bonus points
			IntelligenceValue: 19,   // 3 bonus points (19-10)/3*level
			HP:                tcHP, // 8 HP - 5 damage, but heal them
			MP:                mp.New(character.BaseMP),
		}, true, pts.New(8), pts.New(8), pts.New(11), pts.New(11)},
	}

	for _, tc := range tcs {
		tc.player.RefreshStats(tc.heal)
		if *tc.player.HP.Current != *tc.expectedHP {
			t.Errorf("Expected %d HP, got %d", *tc.expectedHP, *tc.player.HP.Current)
		}
		if *tc.player.MP.Current != *tc.expectedMP {
			t.Errorf("Expected %d MP, got %d", *tc.expectedMP, *tc.player.MP.Current)
		}
		if *tc.player.HP.Max != *tc.expectedMaxHP {
			t.Errorf("Expected %d max HP, got %d", *tc.expectedMaxHP, *tc.player.HP.Max)
		}
		if *tc.player.MP.Max != *tc.expectedMaxMP {
			t.Errorf("Expected %d max MP, got %d", *tc.expectedMaxMP, *tc.player.MP.Max)
		}
	}
}

func TestAddObjectToInventory(t *testing.T) {
	player := &character.Stats{
		Inventory: make([]*object.Object, 0),
	}
	huntingKnife := &object.Object{
		Name:     "Hunting knife",
		Category: "Weapon",
		Weight:   200,
		Stats:    []object.ObjectStat{},
	}

	for _, obj := range player.Inventory {
		if obj == huntingKnife {
			t.Error("Object is already in inventory before adding.")
		}
	}

	player.AddObjectToInventory(huntingKnife, false)

	found := false
	for _, obj := range player.Inventory {
		if obj == huntingKnife {
			found = true
			break
		}
	}

	if !found {
		t.Error("Object was not added to the inventory.")
	}
	if !huntingKnife.InInventory {
		t.Error("Object's InInventory flag was not set to true after adding.")
	}
}

func TestRemoveObjectFromInventory(t *testing.T) {
	player := &character.Stats{
		Inventory: make([]*object.Object, 0),
	}
	huntingKnife := &object.Object{
		Name:     "Hunting knife",
		Category: "Weapon",
		Weight:   200,
		Stats:    []object.ObjectStat{},
	}

	// add to inventory then remove it
	player.AddObjectToInventory(huntingKnife, false)
	for objIndex, obj := range player.Inventory {
		if obj == huntingKnife {
			player.RemoveObjectFromInventory(objIndex)
			break
		}
	}

	if huntingKnife.InInventory {
		t.Error("Object is InInventory and shouldn't.")
	}
}

func TestEquipItem(t *testing.T) {
	player := &character.Stats{
		Inventory: make([]*object.Object, 0),
	}
	huntingKnife := &object.Object{
		Name:     "Hunting knife",
		Category: "Weapon",
		Weight:   200,
		Stats:    []object.ObjectStat{},
	}

	// add to inventory + equip
	player.AddObjectToInventory(huntingKnife, true)

	for objIndex, obj := range player.Inventory {
		if obj == huntingKnife {
			if !player.Inventory[objIndex].Equipped {
				t.Errorf("knife should be equipped and isn't")
			}
			break
		}
	}
}

func TestUnequipItem(t *testing.T) {
	player := &character.Stats{
		Inventory: make([]*object.Object, 0),
	}
	huntingKnife := &object.Object{
		Name:     "Hunting knife",
		Category: "Weapon",
		Weight:   200,
		Stats:    []object.ObjectStat{},
	}

	// add to inventory then equip, then unequip
	player.AddObjectToInventory(huntingKnife, true)
	player.UnequipItem(huntingKnife)

	for objIndex, obj := range player.Inventory {
		if obj == huntingKnife {
			if player.Inventory[objIndex].Equipped {
				t.Errorf("knife shouldn't be equipped and is")
			}
			break
		}
	}
}

func TestComputeWeight(t *testing.T) {
	player := &character.Stats{
		Inventory: make([]*object.Object, 0),
	}
	huntingKnife := &object.Object{
		Name:     "Hunting knife",
		Category: "Weapon",
		Weight:   200,
		Stats:    []object.ObjectStat{},
	}

	player.ChangeGold(100)

	// add to inventory then equip
	player.AddObjectToInventory(huntingKnife, true)

	if player.InventoryWeight != 300 {
		t.Errorf("inventory weight is off, expected 300, got %d", player.InventoryWeight)
	}
	if player.EquippedWeight != 200 {
		t.Errorf("equipped weight is off, expected 200, got %d", player.EquippedWeight)
	}
}

func TestGetGender(t *testing.T) {
	player := &character.Stats{}
	tcs := []struct {
		index          int
		expectedGender string
	}{
		{1, "Female"},
		{2, "Non Binary"},
		{3, "Male"},
	}

	for _, tc := range tcs {
		player.GetGender(tc.index)
		if player.GenderValue != tc.expectedGender {
			t.Errorf("Expected %s gender, got %s", tc.expectedGender, player.GenderValue)
		}
	}
}
