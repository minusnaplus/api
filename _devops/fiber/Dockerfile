# Here we are getting our fastify as Base image
# https://github.com/afteracademy/nodejs-backend-architecture-typescript
# https://github.com/geshan/express-app/blob/main/Dockerfile - better!
# https://github.com/gravityblast/fresh
FROM golang:1.19.3-alpine3.17 as base

# Set the Current Working Directory inside the container
WORKDIR /usr/app
COPY ./erdos/ ./
COPY .air.toml ./
#COPY go.mod go.sum main_compile.go .air.toml ./
RUN ls -la ./*
RUN go mod download
#RUN go mod verify
EXPOSE 3001
# Build the Go app
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o main.bin fiber/main.go

#ENV NODE_ENV=production
FROM scratch as production
COPY --from=base /usr/app/main.bin /main.bin
ENTRYPOINT ["./main.bin"]


FROM base as dev
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]

