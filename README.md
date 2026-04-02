# cli-template-go

Production CLI template using Cobra and Viper.

## Quick Start

```bash
go build -o cli-template
./cli-template hello
go test ./...
```

## Commands

- `hello` - Example hello command

## Development

```bash
make build          # Build binary
make test           # Run tests
make lint           # Lint code
make clean          # Clean build artifacts
```

## Cross-Compilation

```bash
GOOS=linux GOARCH=amd64 go build -o cli-template-linux
GOOS=darwin GOARCH=arm64 go build -o cli-template-macos-arm64
GOOS=windows GOARCH=amd64 go build -o cli-template.exe
```

## License

MIT
