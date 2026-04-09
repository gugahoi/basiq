# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

basiq is a Go CLI tool for interacting with the Basiq API (banking/financial data). It wraps webhook and event APIs into a `urfave/cli/v2` command-line interface.

## Build & Run

```bash
make build        # Compiles binary to ./basiq
make run          # Builds and runs
go build -o basiq .  # Direct build
```

No test suite exists yet. Release builds use goreleaser (`.goreleaser.yaml`).

## Required Environment

*   `BASIQ_APIKEY` — bearer token auth; required for all commands. The CLI exchanges it for a session token via POST `/token`.

## Architecture

**Entry point**: `main.go` creates a `cli.App` with two top-level commands: `webhooks` and `events`.

**Command structure**: Each feature lives in `pkg/<feature>/` with a `root.go` registering subcommands. Each subcommand is a separate package under `pkg/<feature>/<action>cmd/` (e.g., `pkg/webhooks/createcmd/`). Every command package exports `New() *cli.Command` and has a private `exec()` function.

**Client injection**: Root commands (`pkg/webhooks/root.go`, `pkg/events/root.go`) use a `Before` hook to create an authenticated API client via `tools.CreateClient()` or `tools.CreateEventsClient()`, then store it in `cli.Context.App.Metadata["client"]`. Subcommands retrieve it from there.

**API clients**:

* `internal/api/webhooks.gen.go` — **generated** from OpenAPI spec at `.api/webhooks.json` using `oapi-codegen`. Do not edit by hand.
* `internal/api/events/events.go` — **manually written** event client following similar patterns.

**Auth flow** (`tools/client.go`): API key → basic auth POST to `/token` → bearer token used for all subsequent requests. Server URL is hardcoded to `https://au-api.basiq.io/`.

## Conventions

*   Subcommand packages follow `<action>cmd` naming (createcmd, deletecmd, getcmd, listcmd, etc.)
*   Webhook commands output raw pretty-printed JSON; event listing commands use `tabwriter` formatted output
*   Errors are returned (not panicked) from `exec()` functions; `urfave/cli` handles display
*   `mapstructure.Decode()` is used for type conversions from API response maps
*   All methods must have comment with a description about them, in the idiomatic go format
