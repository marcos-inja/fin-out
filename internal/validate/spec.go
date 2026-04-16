package validate

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"

	"fin-out/internal/finance"
	"gopkg.in/yaml.v3"
)

type specRoot struct {
	Version     int                 `yaml:"version"`
	Name        string              `yaml:"name"`
	Formulas    map[string]string   `yaml:"formulas"`
	Output      map[string][]string `yaml:"output"`
	Constraints []string            `yaml:"constraints"`
}

// RunTests runs go test ./... from moduleRoot.
func RunTests(moduleRoot string) error {
	cmd := exec.Command("go", "test", "./...", "-count=1", "-cover", "-coverprofile=coverage.out")
	cmd.Dir = moduleRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go test: %w", err)
	}
	return nil
}

// ValidateSpecFile reads specs/financial_meta.yaml and checks required fields.
func ValidateSpecFile(moduleRoot string) error {
	data, err := os.ReadFile(filepath.Join(moduleRoot, "specs", "financial_meta.yaml"))
	if err != nil {
		return fmt.Errorf("read spec: %w", err)
	}
	var root specRoot
	if err := yaml.Unmarshal(data, &root); err != nil {
		return fmt.Errorf("yaml spec: %w", err)
	}
	if root.Version < 1 || root.Name == "" {
		return fmt.Errorf("invalid spec: version/name")
	}
	if root.Formulas == nil {
		return fmt.Errorf("spec: missing formulas")
	}
	f := root.Formulas["monthly_rate"]
	if f == "" {
		return fmt.Errorf("spec: missing monthly_rate")
	}
	def := root.Output["default"]
	if len(def) < 4 {
		return fmt.Errorf("spec: incomplete default output")
	}
	min := root.Output["minimal"]
	if len(min) < 1 {
		return fmt.Errorf("spec: incomplete minimal output")
	}
	return nil
}

// ValidateVectors checks numeric consistency with the core.
func ValidateVectors() error {
	p0 := finance.Params{Deposit: 100, AnnualRate: 0, Months: 3}
	s0 := finance.BuildSchedule(p0)
	if len(s0.Rows) != 3 || math.Abs(s0.FinalBalance-300) > 1e-9 {
		return fmt.Errorf("vector rate=0 inconsistent")
	}
	p1 := finance.Params{Deposit: 500, AnnualRate: 0.12, Months: 0}
	s1 := finance.BuildSchedule(p1)
	if len(s1.Rows) != 0 || s1.FinalBalance != 0 {
		return fmt.Errorf("vector months=0 inconsistent")
	}
	rm := finance.MonthlyRate(0.14)
	want := math.Pow(1.14, 1.0/12.0) - 1
	if math.Abs(rm-want) > 1e-12 {
		return fmt.Errorf("monthly rate diverges from expected")
	}
	p2 := finance.Params{Deposit: 1e6, AnnualRate: 0.2, Months: 120}
	s2 := finance.BuildSchedule(p2)
	if s2.FinalBalance <= 0 || len(s2.Rows) != 120 {
		return fmt.Errorf("vector extreme values inconsistent")
	}
	a := finance.BuildSchedule(finance.Params{Deposit: 100, AnnualRate: 0.1, Months: 24})
	b := finance.BuildSchedule(finance.Params{Deposit: 100, AnnualRate: 0.1, Months: 24})
	if len(a.Rows) != len(b.Rows) {
		return fmt.Errorf("determinism: sizes diverge")
	}
	for i := range a.Rows {
		if math.Abs(a.Rows[i].Balance-b.Rows[i].Balance) > 1e-12 {
			return fmt.Errorf("determinism: balance diverges")
		}
	}
	return nil
}

// RunAll runs tests, spec validation and vectors.
func RunAll(moduleRoot string) error {
	if err := RunTests(moduleRoot); err != nil {
		return err
	}
	if err := ValidateSpecFile(moduleRoot); err != nil {
		return err
	}
	if err := ValidateVectors(); err != nil {
		return err
	}
	return nil
}
