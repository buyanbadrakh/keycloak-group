FROM golang:1.19-alpine AS go_build

# Linux packages required for build
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN apk add --update tzdata
ENV TZ=Asia/Ulaanbaatar

WORKDIR /app/
COPY . .

# ENV variables
# ENV PUBLIC_KEY="(Keycloak Realm RS256 Public Key)"
# ENV REALM="(Keycloak Realm)"
# ENV URL="https://sso.example.com (keycloak admin URL)"
# ENV GRANT_TYPE="client_credentials"
# ENV CLIENT_ID="(Keycloak Client ID: admin-cli)"
# ENV CLIENT_SECRET="(Keycloak Client Secret)"

ENV APPLICATION_NAME=keycloak-api
ENV APPLICATION_ENVIRONMENT=production
ENV APPLICATION_VERSION=1.0
ENV APPLICATION_BASE_PATH=/api/v1

ENV SERVER_LISTEN=:6060
ENV SERVER_READ_TIMEOUT=60s
ENV SERVER_WRITE_TIMEOUT=60s
ENV SERVER_MAX_CONN_PER_IP=500000
ENV SERVER_MAX_REQUESTS_PER_CONN=500000
ENV SERVER_MAX_KEEP_ALIVE_DURATION=60s

########################
# Create log directory #
########################
RUN echo "=> Adding log directory" && mkdir -p /var/log/keycloak-group

# Install dependencies
# RUN go mod download
RUN go get -d -v

# Stage 2: Build
FROM go_build AS build
# Copy whole source tree
COPY . .
# Build bi project
RUN go install
RUN go build -o keycloak-api

# Stage 3: Build fresh container image without golang
FROM alpine:3.7
RUN apk add --no-cache tzdata ca-certificates
WORKDIR /app/
ENV TZ Asia/Ulaanbaatar

COPY --from=build /app/ /app/

EXPOSE 6060
ENTRYPOINT ["/app/keycloak-api"]