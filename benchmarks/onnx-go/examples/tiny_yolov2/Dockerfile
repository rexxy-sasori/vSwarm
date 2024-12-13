# Stage 1: Build Stage
FROM golang:1.21 as builder

# Set working directory
WORKDIR /app

# Download the repository
RUN git clone https://github.com/rexxy-sasori/onnx-go.git

# Set the working directory to the repository
WORKDIR /app/onnx-go

# Install dependencies using the Makefile
RUN make go-get

# Navigate to the example directory
WORKDIR /app/onnx-go/examples/tiny_yolov2

COPY main.go ./

# Build the static binary for Linux
RUN go get github.com/gin-gonic/gin && go get -u go4.org/unsafe/assume-no-moving-gc && CGO_ENABLED=0 GOOS=linux go build -o image_segmentation

# Stage 2: Create minimal final image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the static binary from the build stage
COPY --from=builder /app/onnx-go/examples/tiny_yolov2/image_segmentation ./
COPY tinyyolov2-7.onnx ./
COPY data/dog.jpg ./
COPY main.go ./

EXPOSE 8080

# Set the binary as the entry point
ENTRYPOINT ["./image_segmentation","-img", "dog.jpg","-model", "tinyyolov2-7.onnx"]

