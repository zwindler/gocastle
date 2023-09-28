package npc

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zwindler/gocastle/pkg/avatar"
	"github.com/zwindler/gocastle/pkg/coord"
	"github.com/zwindler/gocastle/pkg/hp"
	"github.com/zwindler/gocastle/pkg/mp"
)

type Stats struct {
	Name     string
	Pronoun  string
	Dialog   string
	Hostile  bool
	Avatar   avatar.Avatar
	HP       hp.HP
	MP       *mp.MP
	LootXP   int
	LootGold int
}

// New create new NPC.
func New(npc Stats) Stats {
	// Now this function is useless, but it will be useful in the future
	return npc
}

// Spawn creates a copy of a given NPC on given coordinates.
func Spawn(npc Stats, coord coord.Coord) *Stats {
	avatar := avatar.Spawn(npc.Avatar, coord)
	return &Stats{
		Name:    npc.Name,
		Pronoun: npc.Pronoun,
		Avatar:  avatar,
		Dialog:  npc.Dialog,
		Hostile: npc.Hostile,
		HP:      npc.HP,
		MP: func() *mp.MP {
			if npc.MP == nil {
				return npc.MP
			}
			return mp.New(0)
		}(),
		LootXP:   npc.LootXP,
		LootGold: randomizeGoldLoot(npc.LootGold),
	}
}

// HandleNPCDamage returns strings for having nice logs during combat with NPCs.
func (npc *Stats) HandleNPCDamage() string {
	var additionalInfo string

	// Here there are levels of injury
	// I want to give player additional information, but not every time!
	// only when NPC are going from above 80% live to under 80%, for example
	switch {
	case npc.HP.IsDead():
		additionalInfo += fmt.Sprintf("%s is dead.", npc.Name)
	// remaininghp between 80% and 100%
	case npc.HP.IsAlive() && npc.HP.BeetwenPercent(0.8):
		additionalInfo += fmt.Sprintf("%s looks barely injured.", npc.Name)
	// remaininghp between 50% and 80%
	case npc.HP.IsAlive() && npc.HP.BeetwenPercent(0.5):
		additionalInfo += fmt.Sprintf("%s looks injured.", npc.Name)
	// remaininghp between 20% and 50%
	case npc.HP.IsAlive() && npc.HP.BeetwenPercent(0.2):
		additionalInfo += fmt.Sprintf("%s looks seriously injured.", npc.Name)
	// remaininghp between 0% and 20%
	case npc.HP.IsAlive() && npc.HP.Percent() > 0 && npc.HP.Current.Get() < npc.HP.Max.Get() && npc.HP.Current.Get() == npc.HP.Max.Get():
		additionalInfo += fmt.Sprintf("%s looks barely alive.", npc.Name)
	}
	return fmt.Sprintf("you strike at the %s, %s's hit! %s", npc.Name, npc.Pronoun, additionalInfo)
}

// IsNPCDead checks if NPC's HP <= 0.
func (npc *Stats) IsNPCDead() bool {
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
