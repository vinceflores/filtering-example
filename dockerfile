FROM golang:1.22.2

WORKDIR /usr/src/app

# COPY . .
# RUN go mod tidy 

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
