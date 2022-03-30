FROM golang:1.15.12-alpine3.13 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    MYSQL_HOST=0.0.0.0 \
    MYSQL_PORT=3306 \
    MYSQL_USER=root \
    MYSQL_PASSWORD=p455word \
    MYSQL_DBNAME=pos_db  
LABEL maintainer="Ilham Syahidi <ilhamsyahidi66@gmail.com>"
COPY go.mod go.sum /go/src/pos-is-backend/
WORKDIR /go/src/pos-is-backend
RUN go mod download
COPY . /go/src/pos-is-backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/pos-is-backend pos-is-backend

FROM alpine:3.13
# RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /app
COPY --from=builder /go/src/pos-is-backend/build/pos-is-backend .
COPY --from=builder /go/src/pos-is-backend/.env .
COPY --from=builder /go/src/pos-is-backend/static/index.html ./static/
EXPOSE 3030
ENTRYPOINT ["/app/pos-is-backend"]




