// hp is a package that provides a health point system with a maximum and current value.
package hp

import "fmt"

// HP represents a health point system with a maximum and current value.
type HP struct {
	Max     int
	Current int
}

// IsAlive returns true if the current HP is greater than 0.
func (hp *HP) IsAlive() bool {
	return hp.Current > 0
}

// IsDead returns true if the current HP is less than or equal to 0.
func (hp *HP) IsDead() bool {
	return !hp.IsAlive()
}

// Heal adds the given amount to the current HP and caps it at the maximum HP.
func (hp *HP) Heal(amount int) {
	hp.Current += amount
	if hp.Current > hp.Max {
		hp.Current = hp.Max
	}
}

// Damage subtracts the given amount from the current HP and caps it at 0.
func (hp *HP) Damage(amount int) {
	hp.Current -= amount
	if hp.Current < 0 {
		hp.Current = 0
	}
}

// Percent returns the current HP as a percentage of the maximum HP.
func (hp *HP) Percent() float64 {
	return float64(hp.Current) / float64(hp.Max)
}

// PercentString returns the current HP as a percentage string.
func (hp *HP) PercentString() string {
	return fmt.Sprintf("%d%%", int(hp.Percent()*100))
}

// String returns the current and maximum HP as a string.
func (hp *HP) String() string {
	return fmt.Sprintf("%d/%d", hp.Current, hp.Max)
}

// Reset sets the current HP to the maximum HP.
func (hp *HP) Reset() {
	hp.Current = hp.Max
}

// AddMax adds the given amount to the maximum HP.
func (hp *HP) AddMax(amount int) {
	hp.Max += amount
}

// AddCurrent adds the given amount to the current HP.
func (hp *HP) AddCurrent(amount int) {
	hp.Current += amount
}

// SetMax sets the maximum HP to the given amount.
func (hp *HP) SetMax(amount int) {
	hp.Max = amount
}

// SetCurrent sets the current HP to the given amount.
func (hp *HP) SetCurrent(amount int) {
	hp.Current = amount
}

// Set sets the maximum and current HP to the given amount.
func (hp *HP) Set(amount int) {
	hp.Max = amount
	hp.Current = amount
}

// New returns a new HP struct with the given max and current values.
func New(max int) *HP {
	return &HP{
		Max:     max,
		Current: max,
	}
}

// GetCurrent returns the current HP value.
func (hp *HP) GetCurrent() int {
	return hp.Current
}

// GetMax returns the max HP value.
func (hp *HP) GetMax() int {
	return hp.Max
}

// Compute returns the computed Max value based on the given level and base value.
// 8 + 4 by level +
// bonus point for every 3 constitution point above 10 every level.
func Compute(level, base, constitution int) int {
	return base + (4 * (level - 1)) + (constitution-10)/3*level
}

// Compute returns the computed Max value based on the given level and base value.
func (hp *HP) Compute(level, base, constitution int) {
	hp.Max = Compute(level, base, constitution)
}
