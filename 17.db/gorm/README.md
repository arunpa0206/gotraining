
1. Run `go build main.go`

2. Run `go run main.go`

After running the commands, you can find server started on `8080` .

3. To test run 


    curl -d '{"ID":2, "Name":bob, "Email":"bob@gmail.com"}' -H "Content-Type: application/json"-X POST http://localhost:8080/users