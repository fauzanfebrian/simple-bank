FROM golang:1.21.1-alpine as build
WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN cp $GOPATH/bin/migrate .

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine
WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/migrate ./migrate
COPY db/migration ./migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]