FROM golang:1.17
RUN mkdir /app
ADD ./account_gateway/server/ /app/
WORKDIR /app
RUN go get -v github.com/cosmtrek/air
RUN go mod vendor
ENTRYPOINT ["air" , "-d"]
