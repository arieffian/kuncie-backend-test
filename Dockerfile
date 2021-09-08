#############################
# builder
#############################
FROM golang:1.16-buster as builder
LABEL maintainer="https://github.com/arieffian"

ARG BUILD_TYPE
ENV TYPE=$BUILD_TYPE

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
ENV GOPATH /go

WORKDIR /build

# Set timezone
RUN echo Asia/Jakarta > /etc/timezone && \
    rm /etc/localtime && \
    ln -snf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata

# Create appuser.
ENV USER=appuser
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN apt-get update && apt-get -y dist-upgrade

RUN apt-get -y install \
    build-essential \
    libssl-dev \
    ca-certificates

RUN apt-get clean && apt-get -y autoremove && \
    rm -rf /tmp/* /var/tmp/* 

COPY . .

RUN go build -ldflags "-s -w" -o /build/kuncie-backend-test.app cmd/main.go

#############################
# runtime
#############################
FROM debian:buster-slim

# Set working directory
WORKDIR /app

# Set timezone
RUN echo Asia/Jakarta > /etc/timezone && \
    rm /etc/localtime && \
    ln -snf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata

# Copy user
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Install dependencies
RUN apt-get update && apt-get -y install \
    libssl-dev \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Cleanup
RUN apt-get clean && apt-get -y autoremove && \
    rm -rf /tmp/* /var/tmp/* 

# Copy executable file
COPY --from=builder --chown=appuser:appuser /build/kuncie-backend-test.app /app/

# Copy grule file
COPY --from=builder --chown=appuser:appuser rules/discounts.grl /app/rules

# Change directory permission
RUN chown -R appuser:appuser /app/

# Use an unprivileged user.
USER appuser:appuser

# Bind host from any ip
EXPOSE 8080

CMD ["sh", "-c", "/app/kuncie-backend-test.app"]
