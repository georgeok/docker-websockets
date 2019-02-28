FROM golang as build-env
MAINTAINER George Oikonomou <giorgos.n.oikonomou@gmail.com>
WORKDIR /root
COPY . .
RUN GOOS=linux go build -v -ldflags="-w -s" -o /root/server
CMD ["/root/server"]
