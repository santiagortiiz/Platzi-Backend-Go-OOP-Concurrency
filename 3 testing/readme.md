# Testing
- cd 1\ tests/                          # Create package
- go mod init                           # Initialize module
- go test                               # Run tests
- go test -cover                        # See coverage
- go test -coverprofile=coverage.out    # Write coverage to a file
- go tool cover -func=coverage.out      # Coverage details
- go tool cover -html=coverage.out      # Coverage details in the browser

# Profiling
- go test -cpuprofile=cpu.out           # Profiling
- go tool pprof cpu.out                 # Start pprof shell
  - top                                 # show the behavior of the system during the test
  - list <function_name>                # Show execution time of each line
  - list Fibonacci
  - web                                 # show a graphic version of the performance measure
  - pdf                                 # export a report of the performance