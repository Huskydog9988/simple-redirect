# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# copy over go files
COPY *.go ./

# build server
RUN go build -o /simple-redirect

ENV PORT 3000

# export port
EXPOSE ${PORT}

ENTRYPOINT ["/simple-redirect"]

# FROM gcr.io/distroless/base-debian10 AS deploy

# ENV PORT 3000

# WORKDIR /

# COPY --from=build /simple-redirect ./simple-redirect

# # export port
# EXPOSE ${PORT}

# # USER nonroot:nonroot

# ENTRYPOINT ["./simple-redirect"]