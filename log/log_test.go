package log

import "testing"

func TestLog(t *testing.T) {
	log, err := New(&Config{Level: "info"})
	if err != nil {
		t.Error(err)
	}
	log.Info("123")
}
