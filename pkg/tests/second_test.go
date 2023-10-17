package tests

import "testing"

func TestSecond(t *testing.T) {
  t.Run("second test", func (t *testing.T) {
    if false {
      t.Log("Ok!")
    }
    t.Fatal("Err!")
  })
}
