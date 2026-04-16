package finance

import (
	"math"
	"testing"
)

func TestBuildSchedule_NegativeMesesClamped(t *testing.T) {
	p := Params{Deposit: 10, AnnualRate: 0.1, Months: -3}
	s := BuildSchedule(p)
	if len(s.Rows) != 0 {
		t.Fatal("negative months must clamp to 0 rows")
	}
}

func TestScheduleAccumulation(t *testing.T) {
	p := Params{Deposit: 100, AnnualRate: 0.12, Months: 2}
	s := BuildSchedule(p)
	if math.Abs(s.DepositsAccumulated-200) > 1e-9 {
		t.Fatalf("deposits accumulated %v", s.DepositsAccumulated)
	}
}
