# LeetCode Solutions in Go

My solutions to [LeetCode](https://leetcode.com/) problems written in Go.

## Project Structure

Each problem is organized into a separate package containing the solution and its tests.

```text
solutions/
├── p0001TwoSum/
│   ├── solution.go
│   └── solution_test.go
└── ...
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests for a specific problem:

```bash
go test ./solutions/p0001TwoSum
```

Run tests with race detection, randomized order, and coverage:

```bash
go test -race -shuffle=on -cover ./...
```

## Code Quality

The project uses:

- **gofmt** and **goimports** — code formatting.
- **golangci-lint** — static analysis.
- **pre-commit** — automated checks before commits.
- **GitHub Actions** — continuous integration.

Run the linter:

```bash
golangci-lint run ./...
```

Format the code:

```bash
golangci-lint fmt
```

Run all pre-commit hooks:

```bash
pre-commit run --all-files
```
