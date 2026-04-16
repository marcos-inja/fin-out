# Design — fin (financial CLI)

## Goal
Go CLI for goal projections with compound interest and a fixed monthly deposit, oriented around agents and automatic validation against `specs/financial_meta.yaml`.

## Architecture
- `cmd/fin`: Cobra, flags, exit codes, module root discovery.
- `internal/finance`: pure math core (monthly rate, monthly series).
- `internal/core`: parameter assembly and aggregated result (no I/O).
- `internal/output`: text rendering (full and minimal).
- `internal/validate`: spec adherence checks and numeric consistency vectors.
- `tests`: integration/table-driven tests for exported packages.
- `specs`: contract; `agents`: harness roles.

## Data flow
CLI flags → `core.ComputeGoal` → `finance.BuildSchedule` → `output.Render` → stdout.

## Conventions
- `--rate` in the CLI is annual percent (e.g. `14` → `0.14` internally).
- No global state; dependencies injected only via parameters.
- `fin validate` runs Go tests, validates `specs/financial_meta.yaml`, and checks numeric vectors.

## Agents
Planner, Builder, Reviewer and Fixer act in isolation; numeric and I/O decisions must reflect only the spec and this document.
