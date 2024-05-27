# BUILD STAGE ############################################################
FROM golang:1.22.1-alpine3.18 AS builder

WORKDIR /app

# Copy dependecies
COPY ./auth/go.mod ./auth/go.sum ./service/
COPY ./pkg/go.mod ./pkg/go.sum ./pkg/

RUN go work init ./service \
    && go work use ./pkg 

# Download dependencies
RUN go mod download -x

# Copy source code
COPY ./auth ./service
COPY ./pkg ./pkg

# Build the Go app
RUN go build -o main ./service

# RUN STAGE ############################################################
FROM alpine:3.18 AS runner

# Install dependencies when container running
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

# Set working directory
WORKDIR /app

# Copy main app from builder stage
COPY --from=builder /app/main .

RUN ls 

EXPOSE 8000

# Command to run the executable
CMD ["./main"]