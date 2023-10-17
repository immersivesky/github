package tests

import "testing"

func TestSecond(t *testing.T) {
	t.Run("second test", func(t *testing.T) {
		t.Log("Second test finished with status OK")
	})
}
