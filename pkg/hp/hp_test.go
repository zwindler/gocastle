package hp_test

import (
	"testing"

	"github.com/zwindler/gocastle/pkg/hp"
)

func TestHP(t *testing.T) {
	h := hp.New(100)

	// Test IsAlive
	if !h.IsAlive() {
		t.Errorf("Expected HP to be alive, but it is dead")
	}

	// Test Damage
	h.Damage(50)
	if h.Current != 50 {
		t.Errorf("Expected HP to be 50, but got %d", h.Current)
	}

	// Test IsAlive after damage
	if !h.IsAlive() {
		t.Errorf("Expected HP to be alive, but it is dead")
	}

	// Test Damage to death
	h.Damage(50)
	if h.IsAlive() {
		t.Errorf("Expected HP to be dead, but it is alive")
	}

	// Test Heal
	h.Heal(50)
	if h.Current != 50 {
		t.Errorf("Expected HP to be 50, but got %d", h.Current)
	}

	// Test Percent
	if h.Percent() != 0.5 {
		t.Errorf("Expected HP percent to be 0.5, but got %f", h.Percent())
	}

	// Test PercentString
	if h.PercentString() != "50%" {
		t.Errorf("Expected HP percent string to be '50%%', but got '%s'", h.PercentString())
	}

	// Test String
	if h.String() != "50/100" {
		t.Errorf("Expected HP string to be '50/100', but got '%s'", h.String())
	}

	// Test Reset
	h.Reset()
	if h.Current != 100 {
		t.Errorf("Expected HP to be 100, but got %d", h.Current)
	}

	// Test AddMax
	h.AddMax(50)
	if h.Max != 150 {
		t.Errorf("Expected Max HP to be 150, but got %d", h.Max)
	}

	// Test AddCurrent
	h.AddCurrent(50)
	if h.Current != 150 {
		t.Errorf("Expected Current HP to be 150, but got %d", h.Current)
	}

	// Test SetMax
	h.SetMax(200)
	if h.Max != 200 {
		t.Errorf("Expected Max HP to be 200, but got %d", h.Max)
	}

	// Test SetCurrent
	h.SetCurrent(75)
	if h.Current != 75 {
		t.Errorf("Expected Current HP to be 75, but got %d", h.Current)
	}

	// Test Set
	h.Set(50)
	if h.Max != 50 {
		t.Errorf("Expected Max HP to be 50, but got %d", h.Max)
	}
	if h.Current != 50 {
		t.Errorf("Expected Current HP to be 50, but got %d", h.Current)
	}

	// Test GetCurrent
	if h.GetCurrent() != 50 {
		t.Errorf("Expected Current HP to be 50, but got %d", h.GetCurrent())
	}

	// Test GetMax
	if h.GetMax() != 50 {
		t.Errorf("Expected Max HP to be 50, but got %d", h.GetMax())
	}

	// Test IsDead
	if h.IsDead() {
		t.Errorf("Expected HP to be alive, but it is dead")
	}

	// Test HealMax
	h.Heal(60)
	if h.Current != 50 {
		t.Errorf("Expected HP to be 50, but got %d", h.Current)
	}

	// Test Damage to death
	h.Damage(60)
	if !h.IsDead() {
		t.Errorf("Expected HP to be dead, but it is alive")
	}

	// Test Compute
	if hp.Compute(1, 8, 10) != 8 {
		t.Errorf("Expected computed Max HP to be 8, but got %d", hp.Compute(1, 8, 10))
	}
	if hp.Compute(2, 8, 10) != 12 {
		t.Errorf("Expected computed Max HP to be 12, but got %d", hp.Compute(2, 8, 10))
	}
	if hp.Compute(3, 8, 10) != 16 {
		t.Errorf("Expected computed Max HP to be 16, but got %d", hp.Compute(3, 8, 10))
	}
	if hp.Compute(1, 8, 13) != 9 {
		t.Errorf("Expected computed Max HP to be 10, but got %d", hp.Compute(1, 8, 13))
	}
	if hp.Compute(2, 8, 13) != 14 {
		t.Errorf("Expected computed Max HP to be 14, but got %d", hp.Compute(2, 8, 13))
	}
	if hp.Compute(3, 8, 13) != 19 {
		t.Errorf("Expected computed Max HP to be 18, but got %d", hp.Compute(3, 8, 13))
	}

	// Test Compute on HP struct
	h.Compute(1, 8, 10)
	if h.Max != 8 {
		t.Errorf("Expected Max HP to be 8, but got %d", h.Max)
	}
}
