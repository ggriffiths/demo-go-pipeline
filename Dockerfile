FROM golang:1.10.2

ADD . /go/src/github.build.ge.com/212472270/demo-go-pipeline

RUN go install github.build.ge.com/212472270/demo-go-pipeline

ENTRYPOINT /go/bin/demo-go-pipeline

EXPOSE 8080
