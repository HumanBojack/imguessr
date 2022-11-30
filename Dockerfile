FROM golang:1.19-alpine as build

WORKDIR /app

COPY . .

ARG CGO_ENABLED=0

RUN go build -o imguessr ./cmd/api

FROM scratch

WORKDIR /app

COPY --from=build /app/imguessr /app/imguessr

CMD [ "./imguessr"]
