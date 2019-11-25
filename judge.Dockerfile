# build stage
FROM golang:latest AS build
WORKDIR /app

# speed up
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

COPY . .
RUN go build -o ./server . && chmod +x ./server

# prod stage
FROM docker:dind AS prod

WORKDIR /app

COPY --from=build /app/server .
COPY --from=build /app/files .
COPY --from=build /app/statics .

RUN chmod +x ./server
RUN mkdir -p /submit && mkdir -p /compile && mkdir -p /output && mkdir -p /case && mkdir -p /result

ENTRYPOINT ["/app/server"]