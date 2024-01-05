# Stage 1: Build Go application
FROM golang:1.19 AS build

RUN apt-get update
RUN apt-get -y install python3
RUN apt-get -y install python3-pip

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY requirements.txt ./
RUN pip3 install --break-system-packages --no-cache-dir -r requirements.txt

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/.

# Set the entry point command to run the compiled Go application
CMD ["./main"]
