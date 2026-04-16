package output

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"fin-out/internal/finance"
)

func fmtNum(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}

const (
	boxTL = '┌'
	boxTM = '┬'
	boxTR = '┐'
	boxML = '├'
	boxMM = '┼'
	boxMR = '┤'
	boxBL = '└'
	boxBM = '┴'
	boxBR = '┘'
	boxH  = '─'
	boxV  = '│'
)

// Render writes the projection in minimal or full mode.
func Render(minimal bool, s finance.Schedule) string {
	if minimal {
		return fmtNum(s.FinalBalance) + "\n"
	}
	return renderTable(s)
}

func renderTable(s finance.Schedule) string {
	headers := []string{"Month", "Deposits accumulated", "Interest (month)", "Balance"}
	headerAlign := []bool{true, true, true, true}
	dataAlign := []bool{false, false, false, false}

	rows := make([][]string, 0, len(s.Rows))
	for _, r := range s.Rows {
		rows = append(rows, []string{
			strconv.Itoa(r.Month),
			fmtNum(r.DepositsAccumulated),
			fmtNum(r.InterestMonth),
			fmtNum(r.Balance),
		})
	}

	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = cellWidth(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if cw := cellWidth(cell); cw > widths[i] {
				widths[i] = cw
			}
		}
	}

	var b strings.Builder
	b.WriteString(horizLine(widths, boxTL, boxTM, boxTR))
	b.WriteByte('\n')
	b.WriteString(dataLine(headers, widths, headerAlign))
	b.WriteByte('\n')
	b.WriteString(horizLine(widths, boxML, boxMM, boxMR))
	b.WriteByte('\n')
	for _, row := range rows {
		b.WriteString(dataLine(row, widths, dataAlign))
		b.WriteByte('\n')
	}
	b.WriteString(horizLine(widths, boxBL, boxBM, boxBR))
	b.WriteByte('\n')
	return b.String()
}

func horizLine(widths []int, left, mid, right rune) string {
	parts := make([]string, len(widths))
	for i, w := range widths {
		parts[i] = strings.Repeat(string(boxH), w+2)
	}
	return string(left) + strings.Join(parts, string(mid)) + string(right)
}

func dataLine(cells []string, widths []int, leftAlign []bool) string {
	var segs []string
	for i, c := range cells {
		inner := padCell(c, widths[i], leftAlign[i])
		segs = append(segs, " "+inner+" ")
	}
	return string(boxV) + strings.Join(segs, string(boxV)) + string(boxV)
}

func cellWidth(s string) int {
	return utf8.RuneCountInString(s)
}

func padCell(s string, w int, left bool) string {
	sw := cellWidth(s)
	if sw >= w {
		return s
	}
	pad := strings.Repeat(" ", w-sw)
	if left {
		return s + pad
	}
	return pad + s
}
