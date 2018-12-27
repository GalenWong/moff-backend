FROM golang:latest
RUN mkdir -p /moff/moff-backend
WORKDIR /moff/moff-backend
ADD . /moff/moff-backend

RUN go build -o main .
CMD ["/moff/moff-backend/main"]