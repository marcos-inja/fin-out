package core

import (
	"math"
	"testing"
)

func TestComputeGoal_PercentRate(t *testing.T) {
	res := ComputeGoal(GoalInput{
		Target:        30000,
		Deposit:       7000,
		AnnualRatePct: 14,
		Months:        5,
		Minimal:       false,
	})
	if math.Abs(res.Params.AnnualRate-0.14) > 1e-12 {
		t.Fatalf("annual rate decimal %v", res.Params.AnnualRate)
	}
	if len(res.Schedule.Rows) != 5 {
		t.Fatalf("rows %d", len(res.Schedule.Rows))
	}
}

func TestComputeGoal_ClampNegatives(t *testing.T) {
	res := ComputeGoal(GoalInput{Deposit: -10, AnnualRatePct: -5, Months: -2, Minimal: true})
	if res.Params.Deposit != 0 || res.Params.AnnualRate != 0 || res.Params.Months != 0 {
		t.Fatalf("clamp %+v", res.Params)
	}
}
