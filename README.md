# cli-template-go

Production-grade CLI template using Cobra and Viper.

## Features

- **Cobra Framework**: Enterprise CLI framework used by major Go projects
- **Viper Config**: Powerful configuration management (env, files, flags)
- **Testing**: testify for robust unit testing
- **Cross-Compilation**: Multi-platform binary building (Windows, macOS, Linux)
- **Code Quality**: golangci-lint pre-configured
- **CI/CD**: GitHub Actions for testing on multiple platforms
- **Release**: goreleaser for automated multi-platform releases
- **Documentation**: Complete command guides and examples

## Quick Start

### Prerequisites

- Go 1.21+
- Make (optional but recommended)

### Build and Run

```bash
# Build the CLI
go build -o cli-template

# Run a command
./cli-template hello

# Run tests
go test ./...

# Build optimized release
CGO_ENABLED=0 go build -ldflags="-s -w" -o cli-template

# Cross-compile
GOOS=linux GOARCH=amd64 go build -o cli-template-linux
GOOS=darwin GOARCH=arm64 go build -o cli-template-macos-arm64
GOOS=windows GOARCH=amd64 go build -o cli-template.exe
```

## Project Structure

```
cli-template-go/
├── cmd/
│   ├── root.go              # Root command and global config
│   ├── hello.go             # Hello command
│   ├── config.go            # Config commands
│   └── version.go           # Version command
├── pkg/
│   ├── app/
│   │   └── app.go           # Application logic
│   └── utils/
│       └── helpers.go       # Helper functions
├── internal/
│   └── config/
│       └── config.go        # Configuration structs
├── test/
│   └── integration_test.go  # Integration tests
├── .github/
│   └── workflows/           # GitHub Actions CI/CD
│       ├── test.yml         # Test workflow
│       ├── lint.yml         # Linting workflow
│       └── release.yml      # Release workflow
├── .goreleaser.yaml         # Release configuration
├── go.mod                   # Module definition
├── go.sum                   # Dependency checksums
├── Makefile                 # Build automation
└── README.md                # This file
```

## Available Commands

### `hello`

Simple greeting command.

```bash
go build -o cli-template
./cli-template hello --name "Alice"
# Output: Hello, Alice!
```

### `config`

Configuration management commands.

```bash
./cli-template config set key value
./cli-template config get key
```

### Development Commands (Makefile)

| Command | Purpose |
|---------|---------|
| `make build` | Build the CLI binary |
| `make test` | Run all tests |
| `make lint` | Run golangci-lint |
| `make clean` | Remove build artifacts |

## Testing

Tests use Go's built-in testing framework with testify assertions.

### Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package
go test -v ./cmd

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## CI/CD Workflows

### `test.yml`
- Runs on: Windows, macOS, Linux
- Go versions: 1.21, 1.22
- Coverage reporting

### `lint.yml`
- golangci-lint with multiple linters
- Checks for code quality issues
- Runs on Ubuntu (fastest)

### `release.yml`
- Triggered on version tags (v*.*.*)
- Uses goreleaser for cross-compilation
- Creates release artifacts for all platforms

## Cross-Platform Support

The template includes configurations for multiple OS/architecture combinations:

```go
// Supported builds
- Linux x86_64 (amd64)
- Linux ARM64 (arm64)
- macOS x86_64 (amd64)
- macOS ARM64 (arm64)
- Windows x86_64 (amd64)
- Windows ARM64 (arm64)
```

Build for specific platform:

```bash
GOOS=linux GOARCH=amd64 go build -o cli-template-linux-amd64
GOOS=darwin GOARCH=arm64 go build -o cli-template-macos-arm64
GOOS=windows GOARCH=amd64 go build -o cli-template-windows-amd64.exe
```

## Publishing Releases

### Using goreleaser

1. Install goreleaser:
   ```bash
   brew install goreleaser  # macOS
   # or download from https://goreleaser.com/install/
   ```

2. Create a release:
   ```bash
   git tag -a v0.1.0 -m "Release v0.1.0"
   git push origin v0.1.0
   # GitHub Actions will build and publish automatically
   ```

3. Verify the release:
   ```bash
   gh release view v0.1.0
   ```

## Configuration with Viper

Viper supports multiple config sources (priority order):
1. Command-line flags
2. Environment variables
3. Configuration files
4. Default values

Example:

```go
viper.SetDefault("timeout", 30)
viper.BindEnv("timeout", "CLI_TIMEOUT")
viper.BindPFlag("timeout", cmd.Flags().Lookup("timeout"))

timeout := viper.GetInt("timeout")
```

## Extending the Template

### Adding New Commands

1. Create new file in `cmd/`
2. Follow pattern in `cmd/hello.go`
3. Register in `cmd/root.go` with `rootCmd.AddCommand()`
4. Write tests in `test/`

Example:

```go
// cmd/goodbye.go
var goodbyeCmd = &cobra.Command{
    Use:   "goodbye",
    Short: "Say goodbye",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Goodbye!")
    },
}

func init() {
    rootCmd.AddCommand(goodbyeCmd)
}
```

### Adding Dependencies

```bash
go get github.com/user/package
go mod tidy  # Clean up unused dependencies
```

## Troubleshooting

### Build errors

- Check Go version: `go version` (should be 1.21+)
- Clean build cache: `go clean -cache`
- Verify dependencies: `go mod tidy`

### Test failures

- Run with verbose output: `go test -v ./...`
- Check test file paths match pattern `_test.go`
- Verify test function names start with `Test`

### Cross-compilation issues

- Linux to macOS: May need cgo disabled: `CGO_ENABLED=0`
- Windows issues: Check path separators in code
- ARM64: Ensure dependencies support the architecture

## Best Practices

- Use packages to organize code (cmd/, pkg/, internal/)
- Write comprehensive tests for all public functions
- Document exported functions with comments
- Use error wrapping with context
- Handle flags properly with cobra
- Use Viper for configuration management
- Keep commands focused and composable
- Log important operations

## Performance Considerations

- Compiled Go binaries are fast and lightweight
- Single binary deployment (no runtime dependencies)
- Cross-platform compatibility without Docker
- Efficient concurrent operations with goroutines

Build optimized release:
```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o cli-template
```

- `-s`: Strip symbols
- `-w`: Strip DWARF debug info
- Reduces binary size significantly

## Dependencies

### Core
- **Cobra** - CLI framework
- **Viper** - Configuration management

### Testing
- **testify** - Testing assertions and mocks

### Build
- **goreleaser** - Release automation

## License

MIT - See LICENSE file for details

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Write tests: `go test ./...`
5. Run linting: `make lint`
6. Run all tests: `make test`
7. Submit a pull request

## Resources

- [Cobra Documentation](https://cobra.dev)
- [Viper Documentation](https://github.com/spf13/viper)
- [Go by Example](https://gobyexample.com)
- [Effective Go](https://golang.org/doc/effective_go)
- [goreleaser Guide](https://goreleaser.com)
