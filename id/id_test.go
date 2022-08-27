package id

import "testing"

func TestSnowflakeWithLocal(t *testing.T) {
	idGen := NewWithLocal(0)
	t.Log(idGen.Generate().Int64())
	t.Log(idGen.Generate().Int64())
	t.Log(idGen.Generate().Int64())
}
