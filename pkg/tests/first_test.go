package tests

import "testing"

func TestFirst(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		t.Log("First test finished with status OK!")
	})
}
