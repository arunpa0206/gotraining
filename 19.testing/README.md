Write your sample tests looking up on the code and referencing the following

https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go
 Run go test using
 go test

 Run benchmark using
 go test -bench=.

 go test -cover

 go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html