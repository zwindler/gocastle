// mp is a package that provides a mana point system with a maximum and current value.
package mp

import (
	"fmt"

	"github.com/zwindler/gocastle/pkg/pts"
)

// MP represents a mana point system with a maximum and current value.
type MP struct {
	Max     *pts.Point
	Current *pts.Point
}

// Percent returns the current MP as a percentage of the maximum MP.
func (mp *MP) Percent() float64 {
	return float64(mp.Current.Get()) / float64(mp.Max.Get())
}

// PercentString returns the current MP as a percentage string.
func (mp *MP) PercentString() string {
	return fmt.Sprintf("%d%%", int(mp.Percent()*100))
}

// String returns the current and maximum MP as a string.
func (mp *MP) String() string {
	return fmt.Sprintf("%d/%d", mp.Current.Get(), mp.Max.Get())
}

// Reset sets the current MP to the maximum MP.
func (mp *MP) Reset() {
	mp.Current.Set(mp.Max.Get())
}

// Set sets the current MP to the given value.
func (mp *MP) Set(value int) {
	mp.Current.Set(value)
	mp.Max.Set(value)
}

// New returns a new MP struct with the given max and current values.
func New(max int) *MP {
	return &MP{
		Max:     pts.New(max),
		Current: pts.New(max),
	}
}

// Compute returns the computed Max value based on the given level and base value.
// 8 + 4 by level +
// bonus point for every 3 constitution point above 10 every level.
func Compute(level, base, intelligence int) int {
	return base + (4 * (level - 1)) + (intelligence-10)/3*level
}

// Compute returns the computed Max value based on the given level and base value.
func (mp *MP) Compute(level, base, intelligence int) {
	mp.Max.Set(Compute(level, base, intelligence))
}
