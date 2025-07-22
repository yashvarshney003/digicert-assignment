FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh
RUN go test -v ./...


RUN go build -o main ./cmd/main.go

ENV DB_HOST=my_db
ENV DB_PORT=5432
ENV DB_USER=yashv
ENV DB_PASSWORD=admin
ENV DB_NAME=library

EXPOSE 8080

CMD ["./main"]
