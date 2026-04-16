package core

import (
	"fin-out/internal/finance"
)

// GoalInput mirrors CLI flags before percent conversion.
type GoalInput struct {
	Target        float64
	Deposit       float64
	AnnualRatePct float64
	Months        int
	Minimal       bool
}

// GoalResult aggregates schedule and normalized params.
type GoalResult struct {
	Params   finance.Params
	Schedule finance.Schedule
	Minimal  bool
}

// ComputeGoal applies domain rules and delegates to the finance core.
func ComputeGoal(in GoalInput) GoalResult {
	p := finance.Params{
		Target:     in.Target,
		Deposit:    in.Deposit,
		AnnualRate: in.AnnualRatePct / 100.0,
		Months:     in.Months,
	}
	if p.Months < 0 {
		p.Months = 0
	}
	if p.Deposit < 0 {
		p.Deposit = 0
	}
	if p.AnnualRate < 0 {
		p.AnnualRate = 0
	}
	s := finance.BuildSchedule(p)
	return GoalResult{Params: p, Schedule: s, Minimal: in.Minimal}
}
