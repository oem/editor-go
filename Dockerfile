FROM golang:alpine AS builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go install --tags netgo -v -a --ldflags '-w -extldflags "-static"'

FROM scratch
COPY --from=builder /go/bin/replay /

CMD [ "/replay" ]