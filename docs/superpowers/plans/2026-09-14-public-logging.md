# Public Logging Implementation Plan

> For agentic workers: use superpowers:executing-plans for inline execution.

**Goal:** Remove generated schema dependencies from the public logger.

**Architecture:** Keep existing level functions as wrappers around one standard
log output helper. No DB, buffer, retention goroutines or deletion routines.

**Tech Stack:** Go standard library; existing Go/Python verification.

## Tasks

- [ ] Add pkg/logger/logger_test.go: capture standard log output; assert each
  level, formatted message and caller basename; test 100 concurrent calls.
- [ ] Add source boundary test: logger imports no model/GORM and model directory
  contains no Go sources. Run `go test ./pkg/logger` and observe failure.
- [ ] Replace pkg/logger/logger.go with five wrappers and one helper using
  runtime.Caller(2), filepath.Base, fmt.Sprintf and log.Printf. Keep Fatal non-exiting.
- [ ] Remove the explicit tracked pkg/db/model/*.gen.go source list only. Git
  history remains a recovery source. Do not delete any directory recursively.
- [ ] Run `go mod tidy`, inspect resulting dependency changes and preserve
  unrelated dependencies. Update README with removed persistence API and binary caveat.
- [ ] Run `go test -race -count=1 ./...`, `go build ./...`,
  `PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover -s tests -v`,
  and `git diff --check`. Review changes and commit the explicit file list locally.
- [ ] Report local completion and remaining remote/history/binary boundaries.
