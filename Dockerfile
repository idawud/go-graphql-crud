# build stage
FROM golang:1.14-alpine AS build-env
RUN apk add build-base
ADD . /src
RUN cd /src && go build -o server

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/server /app/
ENTRYPOINT ./server