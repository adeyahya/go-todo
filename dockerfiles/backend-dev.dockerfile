FROM golang:alpine

RUN apk update && apk add --no-cache git
RUN go install github.com/mitranim/gow@latest

WORKDIR /app

COPY backend .

RUN go mod tidy

ENV PORT 3001
EXPOSE 3001

CMD [ "gow", "run", "." ]