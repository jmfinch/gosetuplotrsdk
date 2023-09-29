package quote

import (
	"testing"
)

func TestGetQuoteById(t *testing.T) {
	quoteId := "5cd96e05de30eff6ebccebb1"
	quoteDialog := "Hey, stinker, don't go gettingtoo far ahead."
	got, err := GetQuoteByID(quoteId)
	want := quoteDialog

	if err != nil {
		t.Errorf("encountered error %err", err)
	}
	if got.Dialog != want {
		t.Errorf("got %q, wanted %q", got.Dialog, want)
	}
}