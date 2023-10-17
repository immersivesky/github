package tests

import "testing"

func TestFirst(t *testing.T) {
  t.Run("first test", func (t *testing.T) {
    if true {
      t.Log("Ok!")
    }
  })
}
