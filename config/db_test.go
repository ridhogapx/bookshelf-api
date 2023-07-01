package config

import (
	"testing"
)

func TestDB(t *testing.T) {
	m := Msg()
	if m != "Migrating database..." {
		t.Errorf("Failed to migrating!")
	}
}
