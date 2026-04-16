package output

import (
	"strings"
	"testing"

	"fin-out/internal/finance"
)

func TestRender_Minimal(t *testing.T) {
	s := finance.Schedule{FinalBalance: 1234.5}
	got := Render(true, s)
	if strings.TrimSpace(got) != "1234.50" {
		t.Fatalf("got %q", got)
	}
}

func TestRender_Full(t *testing.T) {
	s := finance.Schedule{
		Rows: []finance.MonthRow{
			{Month: 1, DepositsAccumulated: 100, InterestMonth: 0, Balance: 100},
		},
	}
	got := Render(false, s)
	if !strings.Contains(got, "Month") || !strings.Contains(got, "Deposits accumulated") {
		t.Fatalf("missing headers: %q", got)
	}
	if !strings.Contains(got, "100.00") || !strings.Contains(got, "1") {
		t.Fatalf("missing data: %q", got)
	}
	if !strings.Contains(got, "┌") || !strings.Contains(got, "└") {
		t.Fatalf("expected box-drawing table: %q", got)
	}
}

func TestRender_Full_Empty(t *testing.T) {
	got := Render(false, finance.Schedule{})
	if !strings.Contains(got, "Month") {
		t.Fatalf("expected header when empty: %q", got)
	}
}
