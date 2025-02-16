# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile

# ---- Build Stage ----
# FROM golang:1.20 AS builder

# # Set working directory
# WORKDIR /build

# # Copy Go source code
# COPY main.go go.mod go.sum ./

# # Download Go dependencies
# RUN go mod tidy

# # Compile the Go binary
# RUN go build -o redelivery main.go


# ---- Final Home Assistant Image ----
ARG BUILD_FROM
FROM ${BUILD_FROM}

# Execute during the build of the image
# ARG TEMPIO_VERSION BUILD_ARCH
# RUN \
#     curl -sSLf -o /usr/bin/tempio \
#     "https://github.com/home-assistant/tempio/releases/download/${TEMPIO_VERSION}/tempio_${BUILD_ARCH}"

WORKDIR /app

COPY rootfs/ /
RUN chmod +x /etc/services.d/redelivery/run

# Copy compiled Go binary from builder stage
# COPY --from=builder /build/redelivery /app/redelivery

EXPOSE 3140

CMD [ "/etc/services.d/redelivery/run" ]