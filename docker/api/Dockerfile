FROM golang:1.15-alpine

ENV CGO_ENABLED=0

ADD ./_source/api /app/api/_build/

WORKDIR /app/api/_build/

RUN go install
RUN go build -o ./api .

WORKDIR /app/

RUN cp ./api/_build/api ./server_api

CMD [ "./server_api" ]