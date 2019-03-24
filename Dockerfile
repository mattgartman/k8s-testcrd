FROM golang:latest as build
ADD . /go/src/k8s-testcrd
WORKDIR /go/src/k8s-testcrd
RUN go get && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o k8s-testcrd

FROM alpine
WORKDIR /app
COPY --from=build /go/src/k8s-testcrd/k8s-testcrd /app/
CMD ./k8s-testcrd