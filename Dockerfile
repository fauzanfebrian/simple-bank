FROM golang:1.21.1-alpine as build
WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine
WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080
CMD [ "/app/main" ]