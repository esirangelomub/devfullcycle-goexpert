# Tests

## General

```bash
go test .
go test . -v
```

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Benchmarks

```bash
go test -bench=.
```

```bash
go test -bench=. -run=^#
```

```bash
go test -bench=. -benchmem
```

```bash
go test -bench=. -benchmem -cpuprofile=cpu.out
go tool pprof cpu.out
```

```bash
go test -bench=. -benchmem -memprofile=mem.out
go tool pprof mem.out
```

```bash
go test -bench=. -benchmem -blockprofile=block.out
go tool pprof block.out
```

```bash
go test -bench=. -benchmem -mutexprofile=mutex.out
go tool pprof mutex.out
```

```bash
go test -bench=. -benchmem -trace=trace.out
go tool trace trace.out
```

```bash
go test -bench=. -benchmem -trace=trace.out
go tool trace -http=:8080 trace.out
```

## Fuzzing

```bash
go test -fuzz=. -run=^#
```

```bash
go test -fuzz=. -run=^# -fuzztime=10s
```

```bash
go test -fuzz=. -run=^# -fuzztime=10s -fuzzworker=4
```