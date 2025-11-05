# Go (Golang) Mastery Roadmap

> Track your progress. Check off as you learn.  
> Add examples, notes, or links in each section.

## Tier 1: Foundations

> You can write and run basic programs.

- [ ] `package main` — executable vs library packages
- [ ] `func main()` — program entry point
- [ ] `import` statements — bringing in packages
- [ ] `fmt` package — `Println`, `Printf`, `Scan`
- [ ] Variables: `var`, `:=`, zero values
- [ ] Basic types: `int`, `float64`, `bool`, `string`
- [ ] Constants: `const`
- [ ] `go run`, `go build`, `go fmt`

## Tier 2: Control Flow

> You can make decisions and repeat code.

- [ ] `if`, `else if`, `else`
- [ ] `for` loops (only loop in Go!)
- [ ] `switch` (with fallthrough, expressions)
- [ ] `break`, `continue`, labeled statements
- [ ] `defer` — delayed execution

## Tier 3: Functions & Organization

> You write reusable, clean code.

- [ ] Function syntax: parameters, returns, multiple returns
- [ ] Named return values
- [ ] Anonymous functions
- [ ] Closures
- [ ] `init()` functions
- [ ] Package organization (`internal/`, `cmd/`, `pkg/`)
- [ ] `go mod init`, `go.mod`, `go.sum`

## Tier 4: Data Structures

> You store and manage data.

- [ ] Arrays (fixed size)
- [ ] Slices: `make`, `append`, `len`, `cap`, slicing
- [ ] Maps: `make`, literals, `delete`, range
- [ ] Structs: fields, embedding, tags
- [ ] Zero values of composite types

## Tier 5: Pointers & Memory

> You control memory and performance.

- [ ] `*` and `&` — pointers and addresses
- [ ] `new()` vs `&Struct{}`
- [ ] Pointer receivers vs value receivers
- [ ] Escape analysis (understand when things go to heap)
- [ ] Stack vs heap basics

## Tier 6: Methods & Interfaces

> You write flexible, testable code.

- [ ] Methods on structs
- [ ] Value vs pointer receivers
- [ ] Interfaces: implicit satisfaction
- [ ] `interface{}` (empty interface)
- [ ] Type assertions, type switches
- [ ] `Stringer` interface (`fmt.Stringer`)
- [ ] Error handling with `error` interface

## Tier 7: Concurrency
> You write fast, parallel programs.

- [ ] Goroutines: `go func()`
- [ ] `sync.WaitGroup`
- [ ] Channels: buffered vs unbuffered
- [ ] `range` over channels, `close(ch)`
- [ ] `select` statement
- [ ] `sync.Mutex`, `sync.RWMutex`
- [ ] Context: `context.WithCancel`, timeouts
- [ ] Worker pools pattern

## Tier 8: Standard Library Power

> You use Go like a pro.

- [ ] `io`, `ioutil`, `bufio` — reading/writing
- [ ] `os`, `os/exec` — system interaction
- [ ] `encoding/json` — JSON marshal/unmarshal
- [ ] `net/http` — build servers & clients
- [ ] `time`, `time.After`, `time.Ticker`
- [ ] `log` package
- [ ] `testing` — `go test`, table tests, benchmarks
- [ ] `reflect` (use sparingly)

## Tier 9: Tooling & Best Practices

> You write production-ready Go.

- [ ] `go vet`, `staticcheck`, `golint`
- [ ] `go test -cover`, coverage reports
- [ ] Benchmarks: `go test -bench=.`
- [ ] Profiling: `pprof`
- [ ] Error wrapping (`%w`, `errors.Is`, `errors.As`)
- [ ] Custom errors with `errors.New`, `fmt.Errorf`
- [ ] Generics (Go 1.18+) — type parameters
- [ ] Workspace mode (`go work`)

## Tier 10: Mastery Projects

> Apply everything. Build real apps.

- [ ] CLI tool with `cobra` or `urfave/cli`
- [ ] REST API with `net/http` or `Gin`
- [ ] JSON API with proper error handling
- [ ] Concurrent web scraper
- [ ] Database app with `sql` or `gorm`
- [ ] Microservice with gRPC
- [ ] Dockerized Go app

## Resources

- [A Tour of Go](https://tour.golang.org)
- [Go by Example](https://gobyexample.com)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Proverbs](https://go-proverbs.github.io/)
