package test

import (
	"testing"
)

func AssertWalletBalance(t *testing.T, got, want int32) {
	t.Helper()
	if got != want {
		t.Errorf("got incorrect wallet balance, got %d, wanted %d", got, want)
	}
}
