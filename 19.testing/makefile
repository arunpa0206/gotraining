test:
		go test .

		
bench:
		go test -bench=.

cover: 
		go test -cover
		go test -coverprofile=c.out
		go tool cover -html=c.out -o coverage.html
		open coverage.html

