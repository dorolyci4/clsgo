REM go test ./test -v
REM -run=none means only benchmark test executed
REM -benchtime=10s means time to running benchmark
go test ./test -v -bench=. -run=none -memprofile memprofile.out -cpuprofile profile.out