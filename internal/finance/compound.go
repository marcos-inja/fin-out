package finance

import "math"

// Params holds pure numeric inputs (annual_rate as decimal).
type Params struct {
	Target     float64
	Deposit    float64
	AnnualRate float64
	Months     int
}

// MonthRow is one line of the monthly series per spec.
type MonthRow struct {
	Month               int
	DepositsAccumulated float64
	InterestMonth       float64
	Balance             float64
}

// Schedule is the full projection result.
type Schedule struct {
	Rows                []MonthRow
	FinalBalance        float64
	DepositsAccumulated float64
}

// MonthlyRate implements monthly_rate = (1 + annual_rate)^(1/12) - 1.
func MonthlyRate(annualRate float64) float64 {
	return math.Pow(1+annualRate, 1.0/12.0) - 1
}

// BuildSchedule generates the series 1..months with deposit after interest on previous balance.
func BuildSchedule(p Params) Schedule {
	if p.Months < 0 {
		p.Months = 0
	}
	rm := MonthlyRate(p.AnnualRate)
	rows := make([]MonthRow, 0, p.Months)
	var balance float64
	var depositAcc float64

	for m := 1; m <= p.Months; m++ {
		interest := balance * rm
		afterInterest := balance + interest
		newBalance := afterInterest + p.Deposit
		depositAcc += p.Deposit
		rows = append(rows, MonthRow{
			Month:               m,
			DepositsAccumulated: depositAcc,
			InterestMonth:       interest,
			Balance:             newBalance,
		})
		balance = newBalance
	}

	var finalBalance float64
	if len(rows) > 0 {
		finalBalance = rows[len(rows)-1].Balance
	}

	return Schedule{
		Rows:                rows,
		FinalBalance:        finalBalance,
		DepositsAccumulated: depositAcc,
	}
}
