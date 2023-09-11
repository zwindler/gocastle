package pts_test

import (
	"testing"

	"github.com/zwindler/gocastle/pkg/pts"
)

func TestNew(t *testing.T) {
	p := pts.New(10)
	if p == nil {
		t.Errorf("New returned nil")
	}
	if p.Get() != 10 {
		t.Errorf("New did not set the correct value")
	}
}

func TestAdd(t *testing.T) {
	p := pts.New(10)
	p.Add(5)
	if p.Get() != 15 {
		t.Errorf("Add did not add the correct value")
	}
}

func TestSub(t *testing.T) {
	p := pts.New(10)
	p.Sub(5)
	if p.Get() != 5 {
		t.Errorf("Subtract did not subtract the correct value")
	}
}

func TestReset(t *testing.T) {
	p := pts.New(10)
	p.Reset()
	if p.Get() != 0 {
		t.Errorf("Reset did not reset the value to 0")
	}
}

func TestSet(t *testing.T) {
	p := pts.New(10)
	p.Set(20)
	if p.Get() != 20 {
		t.Errorf("Set did not set the correct value")
	}
}

func TestIsZero(t *testing.T) {
	p := pts.New(0)
	if p.IsPositive() {
		t.Errorf("IsZero returned false for a zero value")
	}
	p.Set(10)
	if !p.IsPositive() {
		t.Errorf("IsZero returned true for a non-zero value")
	}
}

func TestString(t *testing.T) {
	p := pts.New(10)
	if p.String() != "10" {
		t.Errorf("String did not return the correct string representation")
	}
}
