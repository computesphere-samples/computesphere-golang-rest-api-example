FROM golang:1.22.3

ENV GIN_MODE=release

WORKDIR /app
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# Copy the source code
COPY . ./
RUN ls -la
# Build
RUN go build -o /computesphere-golang-rest-api-example

EXPOSE 8000
CMD [ "/computesphere-golang-rest-api-example" ]