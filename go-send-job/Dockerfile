# syntax=docker/dockerfile:1

FROM golang:1.19 AS build-stage
WORKDIR /app





COPY go.mod go.sum ./
COPY ./internal ./internal
COPY ./main.go ./main.go
COPY ./configs ./configs

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-send-job

from gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /go-send-job /go-send-job
COPY --from=build-stage ./app/configs /configs

EXPOSE 8082

CMD ["./go-send-job"]