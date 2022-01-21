FROM golang

WORKDIR /usr/local/workspace

COPY ./*.go .

RUN go mod init hello && go mod tidy && go build

EXPOSE 13579

CMD ["./hello"]

