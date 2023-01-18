FROM golang:1.19-alpine

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./cmd cmd
COPY ./internal internal
COPY ./pkg pkg

RUN go build ./cmd/...

COPY ./Procfile .

ENV SECTOR_ID=1 \
PORT=5055

WORKDIR /app

CMD [ "/app/service" ]