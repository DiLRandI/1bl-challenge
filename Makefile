
BENCHOUT="bench.txt"
CPUOUT="cpu.txt"
MEMOUT="mem.txt"

init:
	go install golang.org/x/perf/cmd/...@latest

bench:
	go test -benchmem -timeout 30m -count 10 -a --cpuprofile ./$(CPUOUT) -memprofile ./$(MEMOUT) -run=^$  -bench ^Benchmark 1bl-challange/internal/app  > $(BENCHOUT)
