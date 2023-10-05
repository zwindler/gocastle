// hp is a package that provides a health point system with a maximum and current value.
package hp

import (
	"fmt"

	"github.com/zwindler/gocastle/pkg/pts"
)

// HP represents a health point system with a maximum and current value.
type HP struct {
	Max      *pts.Point
	Current  *pts.Point
	Previous *pts.Point
}

// IsAlive returns true if the current HP is greater than 0.
func (hp *HP) IsAlive() bool {
	return hp.Current.IsPositive()
}

// IsDead returns true if the current HP is less than or equal to 0.
func (hp *HP) IsDead() bool {
	return !hp.Current.IsPositive()
}

// Heal adds the given amount to the current HP and caps it at the maximum HP.
func (hp *HP) Heal(amount int) {
	hp.Current.Add(amount)
	if hp.Current.Get() > hp.Max.Get() {
		hp.Current.Set(hp.Max.Get())
	}
}

// Damage subtracts the given amount from the current HP and caps it at 0.
func (hp *HP) Damage(amount int) {
	hp.Previous.Set(hp.Current.Get())
	hp.Current.Sub(amount)
	if !hp.Current.IsPositive() {
		hp.Current.Reset()
	}
}

// Percent returns the current HP as a percentage of the maximum HP.
func (hp *HP) Percent() float64 {
	return float64(hp.Current.Get()) / float64(hp.Max.Get())
}

// BeetwenPercent returns true if the current HP and previous HP are between the given percentages.
func (hp *HP) BeetwenPercent(percent float64) bool {
	return (float64(hp.Current.Get())/float64(hp.Max.Get())) <= percent && (float64(hp.Previous.Get())/float64(hp.Max.Get())) > percent
}

// PercentString returns the current HP as a percentage string.
func (hp *HP) PercentString() string {
	return fmt.Sprintf("%d%%", int(hp.Percent()*100))
}

// String returns the current and maximum HP as a string.
func (hp *HP) String() string {
	return fmt.Sprintf("%d/%d", hp.Current.Get(), hp.Max.Get())
}

// Reset sets the current HP to the maximum HP.
func (hp *HP) Reset() {
	hp.Current.Set(hp.Max.Get())
}

// Set sets the maximum and current HP to the given amount.
func (hp *HP) Set(amount int) {
	hp.Max.Set(amount)
	hp.Current.Set(amount)
	hp.Previous.Set(amount)
}

// New returns a new HP struct with the given max and current values.
func New(max int) *HP {
	return &HP{
		Max:      pts.New(max),
		Current:  pts.New(max),
		Previous: pts.New(max),
	}
}

// Compute returns the computed Max value based on the given level and base value.
// 8 + 4 by level +
// bonus point for every 3 constitution point above 10 every level.
func Compute(level, base, constitution int) int {
	return base + (4 * (level - 1)) + (constitution-10)/3*level
}

// Compute returns the computed Max value based on the given level and base value.
func (hp *HP) Compute(level, base, constitution int) {
	hp.Max.Set(Compute(level, base, constitution))
}

// Copy creates a new HP with copies of Max, Current, and Previous points.
func (hp *HP) Copy() HP {
	max := hp.Max.Copy()
	cur := hp.Current.Copy()
	pre := hp.Previous.Copy()
	return HP{
		Max:      &max,
		Current:  &cur,
		Previous: &pre,
	}
}
