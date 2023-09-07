package hp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zwindler/gocastle/pkg/hp"
)

func TestHP(t *testing.T) {
	hp := hp.New(100)

	// Test initial values
	assert.Equal(t, 100, hp.Max.Get())
	assert.Equal(t, 100, hp.Current.Get())
	assert.True(t, hp.IsAlive())
	assert.False(t, hp.IsDead())
	assert.Equal(t, "100/100", hp.String())
	assert.Equal(t, "100%", hp.PercentString())

	// Test damage
	hp.Damage(50)
	assert.Equal(t, 50, hp.Current.Get())
	assert.True(t, hp.IsAlive())
	assert.False(t, hp.IsDead())
	assert.Equal(t, "50/100", hp.String())
	assert.Equal(t, "50%", hp.PercentString())

	// Test heal
	hp.Heal(25)
	assert.Equal(t, 75, hp.Current.Get())
	assert.True(t, hp.IsAlive())
	assert.False(t, hp.IsDead())
	assert.Equal(t, "75/100", hp.String())
	assert.Equal(t, "75%", hp.PercentString())

	// Test damage to death
	hp.Damage(75)
	assert.Equal(t, 0, hp.Current.Get())
	assert.False(t, hp.IsAlive())
	assert.True(t, hp.IsDead())
	assert.Equal(t, "0/100", hp.String())
	assert.Equal(t, "0%", hp.PercentString())

	// Test reset
	hp.Reset()
	assert.Equal(t, 100, hp.Current.Get())
	assert.True(t, hp.IsAlive())
	assert.False(t, hp.IsDead())
	assert.Equal(t, "100/100", hp.String())
	assert.Equal(t, "100%", hp.PercentString())

	// Test set max
	hp.Set(200)
	assert.Equal(t, 200, hp.Max.Get())
	assert.Equal(t, 200, hp.Current.Get())
	assert.True(t, hp.IsAlive())
	assert.False(t, hp.IsDead())
	assert.Equal(t, "200/200", hp.String())
	assert.Equal(t, "100%", hp.PercentString())

	// Test Heal to max
	hp.Heal(50)
	assert.Equal(t, 200, hp.Current.Get())
	assert.True(t, hp.IsAlive())
	assert.False(t, hp.IsDead())
	assert.Equal(t, "200/200", hp.String())
	assert.Equal(t, "100%", hp.PercentString())

	// Test compute
	hp.Compute(1, 8, 10)
	assert.Equal(t, 8, hp.Max.Get())

	hp.Compute(2, 8, 10)
	assert.Equal(t, 12, hp.Max.Get())

	hp.Compute(3, 8, 10)
	assert.Equal(t, 16, hp.Max.Get())

	hp.Compute(1, 8, 13)
	assert.Equal(t, 9, hp.Max.Get())

	hp.Compute(2, 8, 13)
	assert.Equal(t, 14, hp.Max.Get())

	hp.Compute(3, 8, 13)
	assert.Equal(t, 19, hp.Max.Get())
}

func TestCompute(t *testing.T) {
	assert.Equal(t, 8, hp.Compute(1, 8, 10))
	assert.Equal(t, 12, hp.Compute(2, 8, 10))
	assert.Equal(t, 16, hp.Compute(3, 8, 10))
	assert.Equal(t, 9, hp.Compute(1, 8, 13))
	assert.Equal(t, 14, hp.Compute(2, 8, 13))
	assert.Equal(t, 19, hp.Compute(3, 8, 13))
}
