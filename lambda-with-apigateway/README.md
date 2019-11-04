Step 1: Create rest-api

aws apigateway create-rest-api --name bookstore

Step 2: Get root path id

aws apigateway get-resources --rest-api-id rest-api-id


Step 3: Create resources

aws apigateway create-resource --rest-api-id rest-api-id --parent-id root-path-id --path-part books

Step 4: Register the HTTP method of ANY

aws apigateway put-method --rest-api-id rest-api-id --resource-id resource-id --http-method ANY --authorization-type NONE

Step 5 - create dynamodb table

aws dynamodb create-table --table-name Books --attribute-definitions AttributeName=ISBN,AttributeType=S --key-schema AttributeName=ISBN,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

Step 6 - Add items to table

aws dynamodb put-item --table-name Book --item '{"ISBN": {"S": "978-1420931693"}, "Title": {"S": "The Republic"}, "Author":  {"S": "Plato"}}'
aws dynamodb put-item --table-name Book --item '{"ISBN": {"S": "978-0486298238"}, "Title": {"S": "Meditations"},  "Author":  {"S": "Marcus Aurelius"}}'

Step 7- build lambda function
env GOOS=linux GOARCH=amd64 go build -o /tmp/main books

Step 8 - zip -j /tmp/main.zip /tmp/main

Step 9a get arn for role

aws iam get-role --role-name role-name

Step 9b- deploy lambda function

aws lambda create-function --function-name books --runtime go1.x --role arn:aws:iam::account-id:role/lambda-books-executor  --handler main --zip-file fileb:///tmp/main.zip


Step10a: aws lambda get-function-configuration --function-name books

Step 10b: Integrate Lambda funtion

aws apigateway put-integration --rest-api-id rest-api-id --resource-id resource-id --http-method ANY --type AWS_PROXY --integration-http-method POST --uri arn:aws:apigateway:us-east-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-2:account-id:function:books/invocations

step 11a get user arn
aws sts get-caller-identity

Step 12: ADD Permission

aws lambda add-permission --function-name books --statement-id a-GUID --action lambda:InvokeFunction --principal apigateway.amazonaws.com --source-arn arn:aws:execute-api:us-east-1:account-id:rest-api-id/*/*/*

Step 13: Download libraries

go get github.com/aws/aws-lambda-go/events



Stdep 14: Invoke api

aws apigateway test-invoke-method --rest-api-id rest-api-id --resource-id resource-id --http-method "GET" --path-with-query-string "/books?isbn=978-1420931693"

step 15 Delete dynamodb table

aws dynamodb delete-table --table-name books

Step 16 Delete rest-api

aws apigateway delete-rest-api --rest-api rest-api

Step 17 aws lambda delete-function --function-name books
