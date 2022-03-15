
FROM golang:alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o summon-app .

RUN apk add curl
RUN curl -sSL https://raw.githubusercontent.com/cyberark/summon/main/install.sh | sh
RUN curl -sSL https://raw.githubusercontent.com/cyberark/summon-conjur/main/install.sh -o install2.sh && sh install2.sh

FROM scratch
COPY --from=builder /app/summon-app /app/
COPY --from=builder /app/secrets.yml /app/
COPY --from=builder /usr/local/bin/summon /app/
COPY --from=builder /usr/local/lib/summon/summon-conjur /app/
EXPOSE 8080
CMD ["/app/summon", "-f", "/app/secrets.yml", "-p", "/app/summon-conjur", "/app/summon-app"]