package singleton

import "testing"

func TestNew(t *testing.T) {
	m1 := NewModel()

	m2 := NewModel()

	if m1.ID != m2.ID {
		t.Error(m1.ID, m2.ID, "singleton not same")
	}
}
