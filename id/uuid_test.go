package id

import "testing"

func TestUUID(t *testing.T) {
	t.Log(UUID())
	t.Log(UUID())
	t.Log(UUID())
	t.Log(UUID())
}

func TestUUID32(t *testing.T) {
	t.Log(UUID32())
	t.Log(UUID32())
	t.Log(UUID32())
}
