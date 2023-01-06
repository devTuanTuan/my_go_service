FROM golang:1.19

WORKDIR /usr/src/app

#Hot reload through Air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .

CMD ["air"]