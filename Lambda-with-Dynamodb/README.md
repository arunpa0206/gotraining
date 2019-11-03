Step1: Create Dynamodb table

aws dynamodb create-table --table-name Books --attribute-definitions AttributeName=ISBN,AttributeType=S --key-schema AttributeName=ISBN,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

Step2: Add items to table

aws dynamodb put-item --table-name Books --item '{"ISBN": {"S": "978-1420931693"}, "Title": {"S": "The Republic"}, "Author":  {"S": "Plato"}}'
aws dynamodb put-item --table-name Books --item '{"ISBN": {"S": "978-0486298238"}, "Title": {"S": "Meditations"},  "Author":  {"S": "Marcus Aurelius"}}'

Step3: Download libraries

go get github.com/aws/aws-sdk-go

Step4: Create a db.go 

touch ~/go/src/books/db.go

Step5: create Main.go

touch ~/go/src/books/mani.go

Step6: Build and zip the binary

env GOOS=linux GOARCH=amd64 go build -o /tmp/main books
zip -j /tmp/main.zip /tmp/main

Step7: Update lambda function

aws lambda update-function-code --function-name books --zip-file fileb:///tmp/main.zip

Step8: Invoke lambda funtion

aws lambda invoke --function-name books /tmp/output.json
cat /tmp/output.json



