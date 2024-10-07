FROM golang:1.22-alpine AS builder

RUN apk add build-base

WORKDIR /src

COPY ./src/go.* ./
RUN go mod download

COPY ./src/ ./

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal -g ./doctors/internal/rest/api.go -o ./doctors/internal/rest -ot "json" 

RUN go build -v -o server ./cmd/doctor-search-engine

FROM alpine:latest AS runtime

COPY --from=builder /src/server /src/server
COPY --from=builder /src/cmd/doctor-search-engine/config.yml /src/config.yml
 
CMD ["/src/server"]
