from golang:1.23.6
WORKDIR /usr/src/app
COPY ./tmp/gin-server /usr/src/app/

CMD ["/usr/src/app/gin-server"]
