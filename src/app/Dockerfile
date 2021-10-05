# Base image
FROM golang:1.17-alpine

# Setting the Work directory
WORKDIR /app

# Installing the Air package
RUN wget --no-check-certificate -q -O - https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copy the source files
COPY . .

# Running the air package
CMD ["air"]

# Setting up air
# docker compose run --rm backend air init