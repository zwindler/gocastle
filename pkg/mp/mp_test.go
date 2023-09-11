package mp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zwindler/gocastle/pkg/mp"
)

func TestMP(t *testing.T) {
	mp := mp.New(100)

	// Test initial values
	assert.Equal(t, 100, mp.Max.Get())
	assert.Equal(t, 100, mp.Current.Get())
	assert.Equal(t, "100/100", mp.String())
	assert.Equal(t, "100%", mp.PercentString())

	// Test reset
	mp.Reset()
	assert.Equal(t, 100, mp.Current.Get())
	assert.Equal(t, "100/100", mp.String())
	assert.Equal(t, "100%", mp.PercentString())

	// Test set max
	mp.Set(200)
	assert.Equal(t, 200, mp.Max.Get())
	assert.Equal(t, 200, mp.Current.Get())
	assert.Equal(t, "200/200", mp.String())
	assert.Equal(t, "100%", mp.PercentString())

	// Test compute
	mp.Compute(1, 8, 10)
	assert.Equal(t, 8, mp.Max.Get())

	mp.Compute(2, 8, 10)
	assert.Equal(t, 12, mp.Max.Get())

	mp.Compute(3, 8, 10)
	assert.Equal(t, 16, mp.Max.Get())

	mp.Compute(1, 8, 13)
	assert.Equal(t, 9, mp.Max.Get())

	mp.Compute(2, 8, 13)
	assert.Equal(t, 14, mp.Max.Get())

	mp.Compute(3, 8, 13)
	assert.Equal(t, 19, mp.Max.Get())
}

func TestCompute(t *testing.T) {
	assert.Equal(t, 8, mp.Compute(1, 8, 10))
	assert.Equal(t, 12, mp.Compute(2, 8, 10))
	assert.Equal(t, 16, mp.Compute(3, 8, 10))
	assert.Equal(t, 9, mp.Compute(1, 8, 13))
	assert.Equal(t, 14, mp.Compute(2, 8, 13))
	assert.Equal(t, 19, mp.Compute(3, 8, 13))
}
