ARG BUILDER_IMAGE=golang:alpine

FROM ${BUILDER_IMAGE} AS build

WORKDIR /usr/src/rest

RUN apk update && \
    apk add --no-cache git ca-certificates tzdata curl && \
    update-ca-certificates

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY go.mod .

ENV GO111MODULE on
RUN go mod download
RUN go mod verify

CMD ["air", "-c", "configs/.air.toml"]