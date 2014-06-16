package wordtrend

import (
	"testing"
	"strings"
)

func compare(t *testing.T, expected, actual interface{}, msg ...string) {
	if expected != actual {
		t.Errorf("[%v] -- value differs. Expected [%v], actual [%v]", msg, expected, actual)
	}
}

func TestSetWords(t *testing.T) {
	compare(t, "These are cats Cat and dogs", strings.Join(SetWords("These are cats Cat and dogs"), " "))
}

func TestGetTrending(t *testing.T) {
	trending := GetTrending()

	compare(t, 1, trending["these"])
	compare(t, 2, trending["cat"])
}

func BenchmarkWordtrend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetWords("These are cats Cat and dogs")
		GetTrending()
	}
}
