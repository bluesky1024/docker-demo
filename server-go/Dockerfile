FROM golang

RUN echo 'start build image'

WORKDIR /usr/local/workspace

COPY ./main.go .

RUN go mod init hello && go build

EXPOSE 13579

CMD ["./hello"]

