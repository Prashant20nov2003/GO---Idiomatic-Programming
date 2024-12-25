package pubadder_test

import (
	"testing"

	"github.com/Prashant20nov2003/ch15/sample_code/pubadder"
)

func TestAddNumbers(t *testing.T) {
	result := pubadder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
