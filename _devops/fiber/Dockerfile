# Here we are getting our fastify as Base image
# https://github.com/afteracademy/nodejs-backend-architecture-typescript
# https://github.com/geshan/express-app/blob/main/Dockerfile - better!
# https://github.com/gravityblast/fresh
FROM golang:1.19.3-alpine3.17 as base

#ENV GOPATH /go/src
#ENV GO111MODULE on

# Set the Current Working Directory inside the container
WORKDIR /usr/app
COPY ./erdos/ ./
COPY .air.toml ./
#COPY go.mod go.sum main_compile.go .air.toml ./
RUN ls -la ./*
RUN go mod download
# how to verify own module?? - skiping ???
#RUN go mod verify

# Build the Go app
#RUN go build -o main_compile main_compile.go
#RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main_compile main_compile.go
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o fiber.bin fiber/fiber.go

#ENV NODE_ENV=production
FROM scratch as production
COPY --from=base /usr/app/fiber.bin /fiber.bin
EXPOSE 4004
ENTRYPOINT ["/fiber.bin"]


FROM base as dev
RUN go install github.com/cosmtrek/air@latest
#CMD ["go", "run", "../src/sixth.go"]
CMD ["air", "-c", ".air.toml"]

