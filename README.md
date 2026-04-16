# fin-out

Financial goal projection CLI in Go.

## What it does

`fin` computes a month-by-month projection using:

- **Monthly compounding** derived from an annual rate: \( (1 + r_{annual})^{1/12} - 1 \)
- **End-of-month deposit**: interest is applied first, then the monthly deposit is added

Outputs are either:

- **Full table**: month, accumulated deposits, interest for the month, and balance
- **Minimal**: only the final balance

The executable spec lives at `specs/financial_meta.yaml` and can be validated locally.

## Install / build

Requires **Go 1.22+**.

```bash
go build -o fin ./cmd/fin
```

Or using the Makefile:

```bash
make build
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

This runs:

- `go test ./...` (with coverage output)
- spec checks against `specs/financial_meta.yaml`
- numeric consistency vectors (e.g. zero rate, zero months, determinism)

## Development

### Formatting and checks

```bash
go fmt ./...
go vet ./...
go test ./... -count=1
```

Or:

```bash
make all
```

### Pre-commit hooks (recommended)

This repository includes a `pre-commit` configuration that runs `go fmt`, `go vet`, and `go test`.

```bash
python -m pip install --user pre-commit
pre-commit install
```

## Release artifacts

GitHub Actions can build cross-platform binaries and attach them to a GitHub Release when you push a tag like `v1.2.3`.

