FROM golang:alpine AS binarybuilder
WORKDIR /app
COPY hello.go .
RUN go build hello.go 

EXPOSE 80
CMD ["/app/hello"]



