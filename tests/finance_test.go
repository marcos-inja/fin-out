package tests

import (
	"math"
	"testing"

	"fin-out/internal/finance"
)

func TestTaxaMensal(t *testing.T) {
	cases := []struct {
		annual float64
		want   float64
	}{
		{0, 0},
		{0.12, math.Pow(1.12, 1.0/12.0) - 1},
		{1.0, math.Pow(2.0, 1.0/12.0) - 1},
	}
	for _, tc := range cases {
		got := finance.MonthlyRate(tc.annual)
		if math.Abs(got-tc.want) > 1e-12 {
			t.Fatalf("MonthlyRate(%v)=%v want %v", tc.annual, got, tc.want)
		}
	}
}

func TestBuildSchedule_Table(t *testing.T) {
	cases := []struct {
		name      string
		p         finance.Params
		wantFinal float64
		wantRows  int
	}{
		{
			name:      "taxa_zero",
			p:         finance.Params{Deposit: 50, AnnualRate: 0, Months: 4},
			wantFinal: 200,
			wantRows:  4,
		},
		{
			name:      "meses_zero",
			p:         finance.Params{Deposit: 999, AnnualRate: 0.5, Months: 0},
			wantFinal: 0,
			wantRows:  0,
		},
		{
			name:      "um_mes",
			p:         finance.Params{Deposit: 1000, AnnualRate: 0, Months: 1},
			wantFinal: 1000,
			wantRows:  1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := finance.BuildSchedule(tc.p)
			if len(s.Rows) != tc.wantRows {
				t.Fatalf("rows=%d want %d", len(s.Rows), tc.wantRows)
			}
			if math.Abs(s.FinalBalance-tc.wantFinal) > 1e-9 {
				t.Fatalf("final=%v want %v", s.FinalBalance, tc.wantFinal)
			}
		})
	}
}

func TestBuildSchedule_CompoundPrecision(t *testing.T) {
	p := finance.Params{Deposit: 7000, AnnualRate: 0.14, Months: 5}
	s := finance.BuildSchedule(p)
	if len(s.Rows) != 5 {
		t.Fatalf("rows %d", len(s.Rows))
	}
	rm := finance.MonthlyRate(0.14)
	var balance float64
	var acc float64
	for m := 1; m <= 5; m++ {
		j := balance * rm
		balance = balance + j + p.Deposit
		acc += p.Deposit
		row := s.Rows[m-1]
		if row.Month != m || math.Abs(row.DepositsAccumulated-acc) > 1e-9 {
			t.Fatalf("m=%d acc got %v want %v", m, row.DepositsAccumulated, acc)
		}
		if math.Abs(row.InterestMonth-j) > 1e-9 {
			t.Fatalf("m=%d interest got %v want %v", m, row.InterestMonth, j)
		}
		if math.Abs(row.Balance-balance) > 1e-9 {
			t.Fatalf("m=%d balance got %v want %v", m, row.Balance, balance)
		}
	}
}

func TestBuildSchedule_Large(t *testing.T) {
	p := finance.Params{Deposit: 1e9, AnnualRate: 0.15, Months: 600}
	s := finance.BuildSchedule(p)
	if len(s.Rows) != 600 || s.FinalBalance <= p.Deposit*float64(p.Months) {
		t.Fatalf("large projection inconsistent")
	}
}
