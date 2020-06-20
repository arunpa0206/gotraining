FROM golang:1.8-alpine
RUN apk update && apk add git && go get github.com/gorilla/mux
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 8000