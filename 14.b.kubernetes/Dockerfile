#add the apline base image
FROM golang:1.8-alpine
#update your linux package manager (apk)
#adding git to package manager and get gorilla mux
RUN apk update && apk add git && go get github.com/gorilla/mux
#make directory app
RUN mkdir /app
#copy the source code from the current directory to the /app
COPY . /app
#set the work directory
WORKDIR /app
#build the go source code
RUN go build -o main .
#run the executable
CMD ["/app/main"]
#expose port
EXPOSE 7000