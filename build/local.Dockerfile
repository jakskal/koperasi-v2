FROM golang:1.21.0-alpine
RUN apk --no-cache add gcc g++ make ca-certificates curl git openssh tzdata
ENV TZ=Asia/Jakarta

WORKDIR /go/src/koperasi
COPY . .

ARG CMD_PATH
ENV CMD_PATH=${CMD_PATH}


RUN go install github.com/githubnemo/CompileDaemon@latest

RUN go mod download
RUN go mod verify

ENTRYPOINT CompileDaemon -exclude-dir=".git" -build="go build -v -o /go/bin/app /go/src/koperasi/${CMD_PATH}/main.go" -command="/go/bin/app"
