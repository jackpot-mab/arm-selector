FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o armselector

EXPOSE 8090

CMD ["./armselector"]