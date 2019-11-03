Step 1-Create folder

cd ~/go/src
mkdir books && cd books
touch main.go

Step 2-Download libraries

go get github.com/aws/aws-lambda-go/lambda

Step 3-Uadate main.go

Step 4-Build main.go file

env GOOS=linux GOARCH=amd64 go build -o /tmp/main books

Step 5-Build binary

zip -j /tmp/main.zip /tmp/main

Step6a : aws iam get-role --role-name role-name

Step 6b-create lambda function


aws lambda create-function --function-name books --runtime go1.x --role arn of role  --handler main --zip-file fileb:///tmp/main.zip

Step 7-Invoke lambda function

aws lambda invoke --function-name books /tmp/output.json
cat /tmp/output.json

aws lambda delete-function --funtion-name books
