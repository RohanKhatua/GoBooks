FROM golang:1.21.0

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build .

EXPOSE 3000

RUN ls

CMD [ "./fiber-jwt" ]