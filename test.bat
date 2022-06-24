@REM go test ./test -v
@REM -run=none means only benchmark test executed
@REM -benchtime=10s means time to running benchmark
@REM go test ./test -v -bench=. -memprofile memprofile.out -cpuprofile profile.out

go test ./test -v -bench=.