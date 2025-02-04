FROM golang:1.23

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o amartha-loan

RUN chmod +x amartha-loan

EXPOSE 8081

CMD [ "./amartha-loan" ]