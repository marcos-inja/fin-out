# fin-out

[![CI](https://img.shields.io/github/actions/workflow/status/marcos-inja/fin-out/ci.yml?branch=main)](../../actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/marcos-inja/fin-out)](../../releases)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8)](https://go.dev/)

Financial goal projection CLI in Go, backed by an executable spec and a deterministic finance core.

## What it does

`fin goal` computes a month-by-month projection with:

- **Monthly compounding** derived from an annual rate: \( (1 + r\_{annual})^{1/12} - 1 \)
- **End-of-month deposit order**: interest is applied to the previous balance first, then the deposit is added

Outputs are either:

- **Full table**: month, accumulated deposits, interest for the month, and balance
- **Minimal**: only the final balance

The executable spec lives at `specs/financial_meta.yaml` and can be validated locally via `fin validate`.

## Quickstart

Requires **Go 1.22+**.

```bash
go build -o fin ./cmd/fin
./fin goal --target 30000 --deposit 7000 --rate 14 --months 5
```

## Commands

### `fin goal`

Inputs:

- `--target` (float): target amount (reference; does not change the math)
- `--deposit` (float): monthly deposit
- `--rate` (float): annual rate in percent (e.g. `14` means `0.14` internally)
- `--months` (int): number of months
- `--minimal` (bool): print only final balance

Outputs:

- Full mode prints a table with columns: `Month`, `Deposits accumulated`, `Interest (month)`, `Balance`
- Minimal mode prints a single number: final balance (2 decimals)

### `fin validate`

Runs:

- `go test ./...` (with coverage output)
- spec checks against `specs/financial_meta.yaml`
- numeric consistency vectors (e.g. zero rate, zero months, determinism)

## Install / build

```bash
go build -o fin ./cmd/fin
```

## Usage

### Full projection table

```bash
./fin goal --target 30000 --deposit 7000 --rate 14 --months 5
```

### Minimal output (final balance only)

```bash
./fin goal --deposit 7000 --rate 14 --months 5 --minimal
```

### Validate (tests + spec + numeric vectors)

```bash
./fin validate
```

## Make targets

```bash
make build
make test
make lint
make validate
make all
```

## Development

### Formatting and checks

```bash
go fmt ./...
go vet ./...
go test ./... -count=1
```

### Pre-commit hooks (recommended)

This repository includes a `pre-commit` configuration that runs `go fmt`, `go vet`, and `go test`.

```bash
python -m pip install --user pre-commit
pre-commit install
```

## Release artifacts

GitHub Actions builds cross-platform binaries and attaches them to a GitHub Release when you push a tag like `v1.2.3`.

Artifacts:

- `fin-linux-amd64`, `fin-linux-arm64`
- `fin-darwin-amd64`, `fin-darwin-arm64`
- `fin-windows-amd64.exe`
