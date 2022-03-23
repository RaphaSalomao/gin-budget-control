FROM golang:1.18
EXPOSE 5000
COPY src /go/src
RUN cd /go/src && go mod download && go mod verify
RUN cd /go/src && rm .env && go build -o application application.go && mv application /go/bin/application
CMD [ "application" ] 