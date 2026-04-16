package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"fin-out/internal/core"
	"fin-out/internal/output"
	"fin-out/internal/validate"
)

func main() {
	root := &cobra.Command{
		Use:           "fin",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	var (
		target, deposit, annualRatePct float64
		months                         int
		minimal                        bool
	)

	goalCmd := &cobra.Command{
		Use:   "goal",
		Short: "Goal projection with compound interest",
		RunE: func(cmd *cobra.Command, args []string) error {
			res := core.ComputeGoal(core.GoalInput{
				Target:        target,
				Deposit:       deposit,
				AnnualRatePct: annualRatePct,
				Months:        months,
				Minimal:       minimal,
			})
			fmt.Print(output.Render(res.Minimal, res.Schedule))
			return nil
		},
	}
	goalCmd.Flags().Float64Var(&target, "target", 0, "target amount for the goal")
	goalCmd.Flags().Float64Var(&deposit, "deposit", 0, "monthly deposit")
	goalCmd.Flags().Float64Var(&annualRatePct, "rate", 0, "annual rate in percent (e.g. 14)")
	goalCmd.Flags().IntVar(&months, "months", 0, "number of months")
	goalCmd.Flags().BoolVar(&minimal, "minimal", false, "print only final balance")

	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Tests, spec validation and consistency vectors",
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			modRoot, err := validate.FindModuleRoot(cwd)
			if err != nil {
				return err
			}
			return validate.RunAll(modRoot)
		},
	}

	root.AddCommand(goalCmd, validateCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
