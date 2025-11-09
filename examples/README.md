# `/examples`

Runnable sample apps showcasing the logger.
One folder per example. Each has its own `main.go`.

## Running
`go run ./examples/<name>`

## Conventions
- Self-contained. No external services.
- Writes to stdout/stderr. No files unless stated.
- Uses standard library only, unless a README in the example says otherwise.
- Flags are optional. Sensible defaults.
